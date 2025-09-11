package routes_tipos_documentos

import (
	"genexis/pos/autoservicios/domain/constants"
	controllers_example "genexis/pos/autoservicios/presentation/api/controllers/example"

	"github.com/gin-gonic/gin"
)

func TiposDocumentosRoutes(api *gin.RouterGroup) {
	TiposDocumentos := api.Group(constants.API_CONFIGURACION)
	{
		TiposDocumentos.GET(constants.API_TIPO_DOCUMENTOS, controllers_example.GetExampleController)
	}
}
