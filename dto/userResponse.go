package dto

type UserResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   uint   `json:"status"`

	// Account can be a uuid string
	Account string `json:"account"`
}
