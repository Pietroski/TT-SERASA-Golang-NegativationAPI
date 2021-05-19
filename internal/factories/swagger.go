package factories

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/Pietroski/TT-SERASA-Golang-NegativationAPI/docs"
)

var (
	Swagger iSwagger = &sSwagger{}
)

type iSwagger interface {
	Generate() *gin.Engine
}

type sSwagger struct {}

func (s *sSwagger) Generate() *gin.Engine {
	// TODO: swagger variables and names should be created dynamically with .env variables

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Technical Test SERASA Golang Negativation API"
	docs.SwaggerInfo.Description = "This is a sample server API server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8008" // "localhost:8010" // "xyz.pietroski.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
