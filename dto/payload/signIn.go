package payload

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
