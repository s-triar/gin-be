package server

import (
	"net/http"
	"net/url"
	"time"

	docs "gin-be/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	docs.SwaggerInfo.BasePath = "/api"
	appRoute := NewRoute(server)
	appRoute.RegisterRoutes()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	server.GET("/openapi", func(ctx *gin.Context) {
		location := url.URL{Path: "/swagger/index.html"}
		ctx.Redirect(http.StatusFound, location.RequestURI())
	})

	return server
}
