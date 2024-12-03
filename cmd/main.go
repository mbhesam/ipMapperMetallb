package main

import (
	"ipMapperApi/docs" // import generated docs
	"ipMapperApi/routes"

	"github.com/gin-contrib/cors" // CORS middleware
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // swagger middleware
)

// @title IP Mapper API
// @version 1.0
// @description This is a simple API to show IP mapping.
// @host localhost:8123
// @BasePath /

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "IP Mapper API"
	docs.SwaggerInfo.Description = "This is a simple API to show IP mapping."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8123"
	docs.SwaggerInfo.BasePath = "/"

	router := gin.Default()
	router.Use(cors.Default()) // Use default CORS settings
	// Serve Swagger UI at /docs
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Redirect from /docs to /docs/index.html
	router.GET("/docs", routes.V1redirect)
	// Serve your API route "/v1/bindings"
	router.GET("/v1/bindings", routes.V1ShowAll)
	// Serve your API route "/v1/bind_ip/:ip"
	router.GET("/v1/bind_ip/:ip", routes.V1ShowPerIP)

	// Start the server on port 8123
	router.Run(":8123")
}
