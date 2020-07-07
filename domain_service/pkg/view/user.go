package view

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
