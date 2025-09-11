package controllers_configuracion

import (
	"genexis/pos/autoservicios/domain/entities"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	"genexis/pos/autoservicios/presentation/container"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConfiguracionPosEsclavoController(c *gin.Context) {
	bodyAny, _ := c.Get("validatedBody")
	body := bodyAny.(*comunes_entidades.ConfiguracionInicialEsclavo)
	data, err := container.GetConfiguracionInicialEsclavoService.Execute(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error al obtener configuración POS Auntoservicio esclavo", err))
		return
	}
	c.JSON(http.StatusOK, data)
}
