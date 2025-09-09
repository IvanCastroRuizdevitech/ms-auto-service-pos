package controllers_example

import (
	"errors"
	"genexis/pos/autoservicios/domain/entities"
	"genexis/pos/autoservicios/presentation/container"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetExampleController(c *gin.Context) {
	log.Println("Recibida solicitud en GetExampleController")

	paramIdStr := c.Param("param")
	param, err := strconv.Atoi(paramIdStr)
	if err != nil {
		log.Printf("Error al obtener parametro de la url: %v", err)
		c.JSON(http.StatusBadRequest, entities.NewErrorServerResponse("Error obtener parametro", errors.New("no se pudo obtener parametro de la URL")))
		return
	}

	processedData, err := container.GetExampleService.Execute(param)
	if err != nil {
		log.Printf("Error al procesar datos: %v", err)
		c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error interno del servidor", errors.New("ups, algo salió mal al obtener la información")))
		return
	}

	log.Printf("datos procesados exitosamente: %v", processedData)
	c.JSON(http.StatusOK, entities.NewSuccessServerResponse("Datos procesados exitosamente", processedData))
}
