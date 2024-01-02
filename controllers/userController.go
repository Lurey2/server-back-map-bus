package controllers

import (
	"fmt"
	"net/http"
	"rutasMap/v2/conf"
	"rutasMap/v2/models"
	"rutasMap/v2/services"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var err error
	var user models.User

	if err = c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, conf.Response{
			Status: conf.Error,
			Code:   http.StatusBadRequest,
			Error:  conf.ErrorType{Message: "Datos de usuario incorrecto", Type: "DATA"},
		})
		return
	}

	user.Rol = "user"
	user.Confirm = true

	if err = services.VerifyEmail(&user); err == nil {
		c.JSON(http.StatusBadRequest, conf.Response{
			Status: conf.Error,
			Code:   http.StatusBadRequest,
			Error:  conf.ErrorType{Message: "Email Registrado", Type: "Email"},
		})
		return
	}

	if err = services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, conf.Response{
			Status: conf.Error,
			Code:   http.StatusInternalServerError,
			Error:  conf.ErrorType{Message: "Error en crear usuario", Type: "INTERNAL"},
		})
		return
	}

	c.JSON(http.StatusOK, conf.Response{
		Data:   true,
		Status: conf.Succes,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.UpdateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {

	var user models.User
	var err error
	var token string

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, conf.Response{
			Status: conf.Error,
			Code:   http.StatusBadRequest,
			Error:  conf.ErrorType{Message: "Email Registrado", Type: "DATA"},
		})
		return
	}

	user, err = services.AuthenticateUser(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, conf.Response{
			Status: conf.Error,
			Code:   http.StatusUnauthorized,
			Error:  conf.ErrorType{Message: err.Error(), Type: "LOGIN"},
		})
		return
	}

	token, err = conf.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, conf.Response{
			Status: conf.Error,
			Code:   http.StatusUnauthorized,
			Error:  conf.ErrorType{Message: "Error Generar Token", Type: "TOKEN"},
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("SESSIONID", token, 15552000, "/", "http://ec2-54-196-14-202.compute-1.amazonaws.com", true, true)
	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   models.User{Email: user.Email, Username: user.Username, Rol: user.Rol, Confirm: user.Confirm},
	})
}

func AuthenticateLog(c *gin.Context) {
	var token string = c.Param("token")

	_, err := services.AuthenticateToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, false)
		return
	}
	c.JSON(http.StatusOK, true)
}

func AuthenticateUser(c *gin.Context) {
	var user models.User

	cook, _ := c.Cookie("SESSIONID")

	data, _ := conf.DecodeToken(cook)

	user.ConvertMapStruct(data.Data)

	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   models.User{Email: user.Email, Username: user.Username, Rol: user.Rol, Confirm: user.Confirm},
	})
}

func LoginGoogleInSignUp(c *gin.Context) {

	fmt.Println("Entro")

	var token conf.TokenLogin
	var user models.User

	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Request": err.Error()})
		return
	}

	fmt.Println(token)

	v := googleAuthIDTokenVerifier.Verifier{}
	aud := "930343312248-5v7hdfgf2ckaalo5hffc5tq76ffvqr3p.apps.googleusercontent.com"
	err := v.VerifyIDToken(token.Credential, []string{
		aud,
	})

	if err != nil {

		c.JSON(http.StatusInternalServerError, conf.Response{
			Status: conf.Error,
			Error:  conf.ErrorType{Message: "Coder User not Found", Type: "AUTHORIZATION"},
		})
		return
	}

	claimSet, _ := googleAuthIDTokenVerifier.Decode(token.Credential)

	if services.VerifiySub(&user, claimSet.Sub) != nil {

		user.Email = claimSet.Email
		user.Username = claimSet.FamilyName
		user.Confirm = true
		user.Rol = "user"
		user.Sub = claimSet.Sub

		services.CreateUser(&user)

	}

	tokenEncript, err := conf.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, conf.Response{
			Status: conf.Error,
			Error:  conf.ErrorType{Message: "Error Generate Token", Type: "TOKEN"},
			Code:   http.StatusBadRequest,
		})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("SESSIONID", tokenEncript, 15552000, "/", "http://ec2-54-196-14-202.compute-1.amazonaws.com", true, true)

	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   models.User{Email: user.Email, Username: user.Username, Rol: user.Rol, Confirm: user.Confirm},
	})
}

func LogoutUser(c *gin.Context) {
	var user models.User

	cook, _ := c.Cookie("SESSIONID")
	data, _ := conf.DecodeToken(cook)
	user.ConvertMapStruct(data.Data)

	if user.ID == 0 {
		c.JSON(http.StatusInternalServerError, conf.Response{
			Code:   http.StatusInternalServerError,
			Error:  conf.ErrorType{Message: "User no identificado", Type: "AUTHORIZATION"},
			Status: conf.Error,
		})
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("SESSIONID", "", -1, "/", "http://ec2-54-196-14-202.compute-1.amazonaws.com", true, true)

	c.JSON(http.StatusOK, conf.Response{
		Status: conf.Succes,
		Data:   true,
	})
}
