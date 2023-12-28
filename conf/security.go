package conf

import (
	custommodels "rutasMap/v2/customModels"
	"rutasMap/v2/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

type TokenLogin struct {
	ClientId   string `json:"clientId"`
	Client_id  string `json:"client_id"`
	Credential string `json:"credential"`
	Select_by  string `json:"select_by"`
}

func GenerateToken(user models.User) (string, error) {

	/*tokenClass := &custommodels.TokenUser{ID: (user.ID), Nombre: user.Username,
	StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), // La fecha de expiración del token es en 1 hora
	}}
	*/
	tokenClass := &custommodels.TokenUser{Data: user,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 6, 0).Unix(), // La fecha de expiración del token es en 1 hora
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClass)

	keySecret := viper.GetString("security.secretKey")
	tokenString, err := token.SignedString([]byte(keySecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func DecodeToken(token string) (custommodels.TokenUser, error) {

	var data custommodels.TokenUser
	_, err := jwt.ParseWithClaims(token, &data, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("security.secretKey")), nil
	})

	return data, err
}
