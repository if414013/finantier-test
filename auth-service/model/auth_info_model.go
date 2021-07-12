package model

type GetTokenData struct {
	Token string `json:"token"`
}

type AuthResult struct {
	Authorized bool `json:"isAuthorized"`
	Message    string `json:"message"`
}

type AuthRequest struct {
	Token string `json:"token"`
}
