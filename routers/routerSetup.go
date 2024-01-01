package routers

import (
	"net/http"
	"rutasMap/v2/controllers"
	"rutasMap/v2/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var pathOrigin = "/api/"

func SetRouter(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://ec2-54-196-14-202.compute-1.amazonaws.com/"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Funciona correctamente")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	user := r.Group(pathOrigin + "/user")

	user.POST("/create", controllers.CreateUser)
	user.POST("/login", controllers.Login)
	user.POST("/signUp", controllers.LoginGoogleInSignUp)
	user.POST("/auten/:token", controllers.AuthenticateLog)
	user.PUT("/update", controllers.UpdateUser)
	user.Use(middleware.Authorization()).POST("/authentification", controllers.AuthenticateUser)
	user.POST("/logout", controllers.LogoutUser)
	route := r.Group(pathOrigin + "/route")

	route.GET("", controllers.GetRoute)
	route.GET("/:ID", controllers.FindByID)
	route.POST("/create", controllers.CreateRoute)
	route.POST("/nearbyRoute", controllers.FindBYNearbyRoute)
	route.POST("/nearbyRouteLtnLng", controllers.FindRouteNearbyLatlng)
	route.GET("/getActive", controllers.GetActiveRoute)
	route.Use(middleware.Authorization())
	route.POST("/getRouteUser", controllers.FindRouteByUserID)
	route.POST("/getRoutePointUser", controllers.FindRoutePointByUser)
	route.PUT("/update", controllers.UpdateRoute)
	route.DELETE("/delete/:ID", controllers.DeleteRoute)

	location := r.Group(pathOrigin + "/location")

	location.GET("", controllers.GetLocation)
	location.GET("/:ID", controllers.FindByIDLocation)
	location.POST("/create", controllers.CreateLocation)
	location.PUT("/update", controllers.UpdateLocation)

}
