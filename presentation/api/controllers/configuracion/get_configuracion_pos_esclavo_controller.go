package controllers_configuracion

import (
	"genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	"genexis/pos/autoservicios/presentation/container"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetConfiguracionPosEsclavoController(c *gin.Context) {
	// Obtener el cuerpo de la petición como JSON
	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 500, "mensaje": "Error al leer el cuerpo de la petición", "error": err.Error()})
		return
	}

	// Convertir el JSON a string para pasarlo a la función almacenada
	jsonParams, err := json.Marshal(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "Error al procesar el JSON de la petición", "error": err.Error()})
		return
	}

	// Obtener la instancia del repositorio de parámetros
	paramRepo, err := container.GetContainer().Get("RecuperarParametrosPos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "Error al obtener el repositorio de parámetros", "error": err.Error()})
		return
	}

	repo := paramRepo.(*infrastructura_repositorios.RecuperarParametrosPos)

	// Llamar a la función almacenada
	query := "SELECT * FROM public.fnc_obtener_configuracion_pos_autoservicio($1)"
	args := []interface{}{string(jsonParams)}

	result, err := repo.Client.Select(query, args)
	if err != nil {
		log.Printf("Error al ejecutar la función almacenada: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "Error al obtener la configuración desde la base de datos", "error": err.Error()})
		return
	}

	// Procesar el resultado de la función almacenada
	if len(result) == 0 || len(result[0]) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "No se obtuvo respuesta de la función almacenada"})
		return
	}

	// La función almacenada devuelve un jsonb, que se mapea a string en Go
	jsonResponseStr, ok := result[0][0].(string)
	if !ok {
		log.Printf("Tipo de dato inesperado de la función almacenada: %T", result[0][0])
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "Formato de respuesta inesperado de la función almacenada"})
		return
	}

	var jsonResponse map[string]interface{}
	if err := json.Unmarshal([]byte(jsonResponseStr), &jsonResponse); err != nil {
		log.Printf("Error al parsear la respuesta JSON de la función almacenada: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "mensaje": "Error al parsear la respuesta JSON de la función almacenada", "error": err.Error()})
		return
	}

	// Devolver la respuesta de la función almacenada directamente
	c.JSON(http.StatusOK, jsonResponse)
}


