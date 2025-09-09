package routes_configuracion

import (
    controllers_configuracion "genexis/pos/autoservicios/presentation/api/controllers/configuracion"
    "github.com/gin-gonic/gin"
)

func ConfiguracionRoutes(api *gin.RouterGroup) {
    api.GET("/configuracion-inicial", controllers_configuracion.GetConfiguracionInicialController)
}
	api.GET(constants.API_POS_MAESTRO, controllers_configuracion.GetConfiguracionPosEsclavoController)


