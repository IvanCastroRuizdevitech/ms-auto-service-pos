package presentation_api_middlewares

import (
	"errors"
	"genexis/pos/autoservicios/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateBodyStruct[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var obj T

		if err := c.ShouldBindJSON(&obj); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				// Para mostrar todas las validaciones fallidas, creamos un string concatenado
				var mensajes string
				for _, fe := range ve {
					mensajes += fe.Field() + ": " + fe.Tag() + "; "
				}
				c.JSON(http.StatusBadRequest, entities.NewErrorServerResponse("Campos inválidos en el body", errors.New(mensajes)))
				c.Abort()
				return
			}

			c.JSON(http.StatusBadRequest, entities.NewErrorServerResponse("JSON mal formado", err))
			c.Abort()
			return
		}

		c.Set("validatedBody", obj)
		c.Next()
	}
}
