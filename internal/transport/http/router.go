package http

import (
	"github.com/abdelrhman-elsbagh/pack-calculator/internal/domain/packcalculator"
	packcalculator2 "github.com/abdelrhman-elsbagh/pack-calculator/internal/transport/http/packcalculator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// SetupRouter sets up the HTTP routes and middleware for the API.
// It wires everything together: services, use cases, and handlers.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set up CORS so the frontend (on localhost:5173) can call the API without issues
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow requests from Vue dev server
		AllowMethods:     []string{"POST"},                  // Only POST requests are allowed
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Create the layers: service -> use case -> handler
	calculatorService := packcalculator.NewCalculatorService()
	useCase := packcalculator.NewUseCase(calculatorService)
	handler := packcalculator2.NewPackHandler(useCase)

	// Define the main route for the pack calculator
	r.POST("/calculate", handler.Calculate)

	r.GET("/docs", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("./internal/docs/docs.html")
	})

	r.GET("/docs/openapi.yaml", func(c *gin.Context) {
		c.Header("Content-Type", "application/yaml")
		c.File("./internal/docs/openapi.yaml")
	})

	return r
}
