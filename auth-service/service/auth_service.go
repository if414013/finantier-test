package service

import "auth-service/model"

type AuthService interface {
	GetJwt() (response *model.GetTokenData)
	Authenticate(authRequest model.AuthRequest) (response *model.AuthResult)
}
