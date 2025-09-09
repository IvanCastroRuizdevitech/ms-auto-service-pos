package routes_example

import (
	"genexis/pos/autoservicios/domain/constants"
	controllers_example "genexis/pos/autoservicios/presentation/api/controllers/example"

	"github.com/gin-gonic/gin"
)

func ExampleRoutes(api *gin.RouterGroup) {
	exampleGroup := api.Group(constants.API_EXAMPLE)
	{
		exampleGroup.GET("/:param", controllers_example.GetExampleController)
	}
}
