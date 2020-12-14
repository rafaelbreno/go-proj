package dto

type UserResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Status uint   `json:"status"`

	// Account can be a uuid string
	Account string `json:"account"`
}
