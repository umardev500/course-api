package model

type LoginRequest struct {
	User string `json:"user" validate:"min=6"`
	Pass string `json:"pass" validate:"min=8"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
