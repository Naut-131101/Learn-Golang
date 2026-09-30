package models

type Account struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Balance  float64 `json:"balance"`
}
