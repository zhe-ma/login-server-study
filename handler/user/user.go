package user

type CeateRequest struct {
	Username string `json: "username"`
	Password string `json: "password"`
}
