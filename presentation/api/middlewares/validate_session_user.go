package presentation_api_middlewares

import (
	"encoding/json"
	"errors"
	"genexis/pos/autoservicios/domain/entities"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("DEV") == "true" {
			log.Println("[MIDDLEWARE] SESION DE USUARIO HABILITADA EN MODO DESARROLLO")
			c.Next()
		} else {
			url := os.Getenv("HTTPS_PROTOCOL") + "://" + os.Getenv("AUTH_SERVER_HOST") + ":" + os.Getenv("AUTH_SERVER_PORT") + "/api/v1/session/ws"
			token := c.GetHeader("Authorization")

			if token == "" {
				c.JSON(http.StatusUnauthorized, entities.NewErrorServerResponse("Error Autorización", errors.New("No se ha proporcionado el token de autenticación")))
				c.Abort()
				return
			}

			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error interno", errors.New("Error creando la solicitud")))
				c.Abort()
				return
			}

			req.Header.Set("Authorization", token)

			resp, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error interno", errors.New("Error conectando con el servidor de autenticación")))
				c.Abort()
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error interno", errors.New("Error leyendo la respuesta del servidor de autenticación")))
				c.Abort()
				return
			}

			var result map[string]interface{}

			err = json.Unmarshal([]byte(body), &result)
			if err != nil {
				c.JSON(http.StatusInternalServerError, entities.NewErrorServerResponse("Error interno", errors.New("Error al decodificar el cuerpo JSON")))
				c.Abort()
				return
			}

			if resp.StatusCode != http.StatusOK {
				c.JSON(http.StatusUnauthorized, entities.NewErrorServerResponse("Error Autorización", errors.New("Sesión no válida")))
				c.Abort()
				return
			}

			c.Next()
		}
	}
}
