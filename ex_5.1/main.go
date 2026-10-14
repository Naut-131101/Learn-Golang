package main

import (
	"fmt"
	"sync"
)

var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

////////////////////////////////////////////////////

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

/////////////////////////////////////////

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

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go sayHi("sayHi function execution", &wg)

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("main function execution")
	// }
	// wg.Wait()

	Crawl("http://golang.org/", 4, fetcher)
}
