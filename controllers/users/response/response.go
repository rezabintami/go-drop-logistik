package response

import (
	"go-drop-logistik/modules/users"
)

type Users struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func FromDomain(userDomain users.Domain) Users {
	return Users{
		ID:    userDomain.ID,
		Name:  userDomain.Name,
		Email: userDomain.Email,
	}
}

func TokenFromDomain(accessToken, refreshToken string) Token {
	return Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
