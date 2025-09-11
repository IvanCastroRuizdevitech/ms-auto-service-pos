package controllers_configuracion

import (
	"genexis/pos/autoservicios/domain/entities"
	"genexis/pos/autoservicios/presentation/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConfiguracionInicialController(c *gin.Context) {
	data, err := container.GetConfiguracionInicialService.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error al obtener configuración", err))
		return
	}
	c.JSON(http.StatusOK, data)
}

