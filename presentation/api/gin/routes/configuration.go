package api_routes

import (
	"genexis/pos/autoservicios/domain/constants"
	routes_configuracion "genexis/pos/autoservicios/presentation/api/gin/routes/configuracion"
	routes_example "genexis/pos/autoservicios/presentation/api/gin/routes/example"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func GinConfig() (*gin.Engine, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, PATCH",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          300 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	// API group
	api := router.Group(constants.API_PATH)

	// Routes
	routes_example.ExampleRoutes(api)
	routes_configuracion.ConfiguracionRoutes(api)

	return router, nil
}
