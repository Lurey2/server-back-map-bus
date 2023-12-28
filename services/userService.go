package services

import (
	"errors"
	"fmt"
	"rutasMap/v2/conf"
	custommodels "rutasMap/v2/customModels"
	"rutasMap/v2/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func GetFindUser() ([]models.User, error) {

	var users []models.User

	db, err := conf.InitDb()

	if err != nil {
		fmt.Println("Error")
		return nil, err
	}

	if err := db.Find(&users).Error; err != nil {
		fmt.Println("Error")
		return nil, err
	}

	return users, err
}

func CreateUser(user *models.User) error {

	db, errD := conf.InitDb()
	fmt.Println(user.Password)
	if user.Password != "" {

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)

		if err != nil {
			return err
		}

	}

	if errD != nil {
		fmt.Println("Error ConnectDb ")
	}

	if errD := db.Create(&user).Error; errD != nil {
		fmt.Println("Error Create ")
	}
	return err
}

func UpdateUser(user *models.User) error {
	db, err := conf.InitDb()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	if err := db.Model(&user).Updates(user).Error; err != nil {
		fmt.Println("Error Create ")
	}
	return err
}

func AuthenticateUser(username string, password string) (models.User, error) {
	var user models.User
	db, err := conf.InitDb()
	if err != nil {
		fmt.Println("Error ConnectDb ")
	}

	err = db.Where("email = ?", username).Find(&user).Error
	if err != nil {
		return user, errors.New("Usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("Contraseña incorrecta")
	}

	return user, nil
}

func AuthenticateToken(token string) (*custommodels.TokenUser, error) {

	tokenConfig, err := jwt.ParseWithClaims(token, &custommodels.TokenUser{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Devolver la clave secreta
		return []byte(viper.GetString("security.secretKey")), nil
	})
	if err != nil || tokenConfig == nil {
		return nil, err
	}

	// Verifica que el token sea válido y tiene los datos necesarios
	if claims, ok := tokenConfig.Claims.(*custommodels.TokenUser); ok && tokenConfig.Valid {
		return claims, nil
	} else {
		return nil, errors.New("Token inválido")
	}

}

func VerifiySub(user *models.User, sub string) error {

	db, err := conf.InitDb()

	if err != nil {
		fmt.Println("Error ConnectDb ")
		return err
	}

	err = db.First(&user, "sub = ?", sub).Error

	return err

}
func VerifyEmail(user *models.User) error {

	db, err := conf.InitDb()

	if err != nil {
		fmt.Println("Error ConnectDb ")
		return err
	}

	err = db.First(&user, "email = ?", user.Email).Error

	return err

}
