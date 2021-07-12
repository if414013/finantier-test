package service

import (
	"auth-service/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

func NewAuthService() AuthService {
	return &authServiceImpl{}
}

type authServiceImpl struct{}

func (s authServiceImpl) GetJwt() (response *model.GetTokenData) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	jwtExpiryInMinutes, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN_MINUTES"))
	claims["authorized"] = true
	claims["client"] = "Guest User"
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtExpiryInMinutes)).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))

	if err != nil {
		fmt.Errorf("Something Went Wrong When Generating JWT : %s", err.Error())
		return nil
	}

	return &model.GetTokenData{
		Token: tokenString,
	}
}

func (s authServiceImpl) Authenticate(authRequest model.AuthRequest) (response *model.AuthResult) {
	token, err := jwt.Parse(authRequest.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error when parsing jwt")
		}
		return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
	})

	if err == nil {
		if token.Valid {
			return &model.AuthResult{
				Authorized: true,
				Message:    "Authentication Successfully",
			}
		}
	}
	fmt.Println(err)
	return &model.AuthResult{
		Authorized: false,
		Message:    "Token didn't exist in header or Token Invalid or already Expired, try to re-fetch it from /token endpoint",
	}
}
