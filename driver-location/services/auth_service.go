package services

import (
	"location-service/utils"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AuthenticateUser(creds Credentials) (string, error) {
	return utils.GenerateToken(utils.User{
		Username:      creds.Username,
		Authenticated: true,
	})

}
