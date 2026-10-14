package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

type fakeResult struct {
	body string
	urls []string
}

func sayHi(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	cache := SafeCache{
		visited: make(map[string]bool),
	}

	var wg sync.WaitGroup
	var crawl func(string, int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		if depth <= 0 {
			return
		}
		if !cache.MarkVisited(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			// Crawl(u, depth-1, fetcher)
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}
	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()
}

type SafeCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (c *SafeCache) MarkVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[url] {
		return false
	}
	c.visited[url] = true
	return true
}

type HTTPFetcher struct {
	Client    *http.Client
	LimitHost string
}

func (f HTTPFetcher) Fetch(urlStr string) (string, []string, error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return "", nil, err
	}

	req.Header.Set("User-Agent", "GoCrawler/1.0")

	resp, err := f.Client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", nil, fmt.Errorf("fetch failed: %s status=%s", urlStr, resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", nil, fmt.Errorf("skip non-html: %s content-type=%s", urlStr, contentType)
	}

	reader := io.LimitReader(resp.Body, 2*1024*1024)

	doc, err := html.Parse(reader)
	if err != nil {
		return "", nil, err
	}

	title := extractTitle(doc)
	links := extractLinks(doc, urlStr, f.LimitHost)

	if title == "" {
		title = resp.Status
	}

	return title, links, nil
}

func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		if n.FirstChild != nil {
			return strings.TrimSpace(n.FirstChild.Data)
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		title := extractTitle(child)
		if title != "" {
			return title
		}
	}

	return ""
}

func extractLinks(n *html.Node, pageURL string, limitHost string) []string {
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		return nil
	}

	var links []string

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}

				link, err := url.Parse(attr.Val)
				if err != nil {
					continue
				}

				absoluteURL := baseURL.ResolveReference(link)

				if absoluteURL.Scheme != "http" &&
					absoluteURL.Scheme != "https" {
					continue
				}

				if limitHost != "" && absoluteURL.Host != limitHost {
					continue
				}

				absoluteURL.Fragment = ""

				links = append(links, absoluteURL.String())
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}

	walk(n)

	return links
}

func main() {
	startURL := "https://vtc.edu.vn/"

	parsedURL, err := url.Parse(startURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fetcher := HTTPFetcher{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		LimitHost: parsedURL.Host,
	}

	Crawl(startURL, 2, fetcher)
}
