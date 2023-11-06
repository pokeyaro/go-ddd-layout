package sys

type LoginData struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
}
