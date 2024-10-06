//go:generate CompileDaemon -command=./first-go-api
package main

import (
	"first-go-api/inits"
	"first-go-api/routes"

	_ "first-go-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

// @title           Exemplo API com Gin
// @version         1.0
// @description     Esta é uma API de exemplo usando Gin e Swagger
// @host            localhost:3010
// @BasePath        /api/v1
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterRoutes(r)

	r.Run()

}
