package main
import (
	_ "github.com/CerWL/swagdemo/docs"
	"github.com/CerWL/swagdemo/user/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User API
// @description User microservice server
// @schemes http https
// @securityDefinitions.basic BasicAuth

func main() {
	r := gin.Default()

	r.Get("/users/", handler.ListUser())
	r.Post("/user/", handlerCreateUser())
	r.Run()
}
