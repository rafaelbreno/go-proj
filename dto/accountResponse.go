package dto

type AccountResponse struct {
	ID      uint   `json:"id"`
	Account string `json:"account"`
	Balance int    `json:"balance"`
}
