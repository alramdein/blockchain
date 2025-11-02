// @title Blockchain API
// @version 1.0
// @description A simple blockchain implementation with transfer and balance tracking
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
package main

import (
	"fmt"
	"log"

	api "github.com/alramdein/blockchain/delivery"
	_ "github.com/alramdein/blockchain/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create HTTP handler
	handler := api.NewHTTPHandler(e)
	handler.RegisterRoutes()

	// Add Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	fmt.Println("🚀 Starting blockchain API server on port 8080...")
	fmt.Println("📚 Swagger documentation available at: http://localhost:8080/swagger/index.html")
	fmt.Println()

	log.Fatal(e.Start(":8080"))
}
