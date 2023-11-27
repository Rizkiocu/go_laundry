package dto

type AuthRequest struct {
	Username string
	Password string
}

type AuthRespone struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
