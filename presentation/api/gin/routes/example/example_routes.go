package routes_example

import (
	"genexis/pos/autoservicios/domain/constants"
	controllers_example "genexis/pos/autoservicios/presentation/api/controllers/example"
	presentation_api_middlewares "genexis/pos/autoservicios/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ExampleRoutes(api *gin.RouterGroup) {
	exampleGroup := api.Group(constants.API_EXAMPLE)
	{
		exampleGroup.GET("/:param", presentation_api_middlewares.SessionMiddleware(), controllers_example.GetExampleController)
	}
}
