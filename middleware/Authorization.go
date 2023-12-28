package middleware

import (
	"net/http"
	"rutasMap/v2/conf"
	custommodels "rutasMap/v2/customModels"
	"rutasMap/v2/models"

	"github.com/gin-gonic/gin"
)

func Authorization(authorize ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var tokenModel custommodels.TokenUser

		cookkie, err := c.Cookie("SESSIONID")

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, conf.Response{
					Status: conf.Error,
					Error:  conf.ErrorType{Message: "Sin Autorizacion", Type: "Authorization"},
					Code:   http.StatusUnauthorized,
				},
			)
			return
		}

		tokenModel, err = conf.DecodeToken(cookkie)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, conf.Response{
					Status: conf.Error,
					Error:  conf.ErrorType{Message: "Sin Autorizacion", Type: "Authorization"},
					Code:   http.StatusUnauthorized,
				},
			)
			return
		}

		user.ConvertMapStruct(tokenModel.Data)
		if len(authorize) == 0 {
			c.Next()
		} else {
			for _, v := range authorize {
				if v == user.Rol {
					c.Next()
					return
				}
			}
		}
	}
}
