package controllers_example

import (
	"genexis/pos/autoservicios/domain/entities"
	"genexis/pos/autoservicios/presentation/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTiposDocumentosController(c *gin.Context) {

	data, err := container.GetConfiguracionInicialService.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error al obtener tipos de documentos", err))
		return
	}
	c.JSON(http.StatusOK, data)
}
