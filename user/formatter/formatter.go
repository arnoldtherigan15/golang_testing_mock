package formatter

type LoginResponse struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
