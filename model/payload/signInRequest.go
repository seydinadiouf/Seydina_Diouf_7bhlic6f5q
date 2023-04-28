package payload

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
