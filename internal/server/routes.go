package server

import (
	"gin-be/internal/controller"
	"gin-be/internal/database"
	"gin-be/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteInterface interface {
	RegisterRoutes()
}

type Route struct {
	route *gin.RouterGroup
}

func NewRoute(route *gin.Engine) *Route {
	return &Route{
		route: route.Group(""),
	}
}

func HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Welcome to POS API"

	c.JSON(http.StatusOK, resp)
}

func healthHandler(c *gin.Context) {
	db := database.New()
	c.JSON(http.StatusOK, db.Health())
}

func (route *Route) RegisterRoutes() {

	route.route.GET("/", HelloWorldHandler)

	route.route.GET("/health", healthHandler)

	api := route.route.Group("api")
	v1 := api.Group("v1")
	// ===AUTH===
	v1_auth := v1.Group("auth")
	v1_auth.POST("/register", middleware.AnonymousOnly(), controller.RegisterControllerMethod)
	v1_auth.POST("/login", middleware.AnonymousOnly(), controller.LoginControllerMethod)
	// v1_auth.POST("/forget-password", AnonymousOnly(), func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "Hello from auth forgot-password",
	// 	})
	// })
	v1_auth.GET("/check_email", middleware.AnonymousOnly(), controller.CheckEmailControllerMethod)
	v1_auth.GET("/check_phone", middleware.AnonymousOnly(), controller.CheckPhoneControllerMethod)
	v1_auth.GET("/user", middleware.AuthJWT(), controller.UserProfileControllerMethod)
	// ===AUTH END===
}
