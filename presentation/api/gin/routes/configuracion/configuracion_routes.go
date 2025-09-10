package routes_configuracion

import (
	"genexis/pos/autoservicios/domain/constants"
	controllers_configuracion "genexis/pos/autoservicios/presentation/api/controllers/configuracion"
	presentation_api_middlewares "genexis/pos/autoservicios/presentation/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfiguracionRoutes(api *gin.RouterGroup) {
	configuracion := api.Group(constants.API_CONFIGURACION) 
	{ 
		configuracion.GET(constants.API_POS_CONFIGURACION, controllers_configuracion.GetConfiguracionInicialController)
		configuracion.POST(constants.API_POS_MAESTRO,  presentation_api_middlewares.ValidateBodyStruct[any](), controllers_configuracion.GetConfiguracionPosEsclavoController)
	}	
}
