package aplication_services_configuracion_esclavo

import (
	"encoding/json"
	"errors"
	aplication_usecases_configuracion_esclavo "genexis/pos/autoservicios/aplication/usecases/configuracion_esclavo"
	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
	"genexis/pos/autoservicios/domain/constants"
	"genexis/pos/autoservicios/domain/entities"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	infrastructura_repositorios "genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	"log"
)

type GetConfiguracionInicialEsclavoService struct {
	UseCase       *aplication_usecases_configuracion_esclavo.GetConfiguracionInicialEsclavoUseCase
	HTTPClient    domain_adapters_clients_http.IClientHttp
	ParametroRepo *infrastructura_repositorios.RecuperarParametrosPos
}

func (s *GetConfiguracionInicialEsclavoService) Execute(body any) (*comunes_entidades.ConfiguracionInicial, error) {

	// 1. Recuperar parámetros
	autoservicioMaestro, err := s.ParametroRepo.Consultar("AUTOSERVICIO_MAESTRO")
	if err != nil {
		return nil, err
	}
	autoservicioIPMaestro, err := s.ParametroRepo.Consultar("AUTOSERVICIO_IP_MAESTRO")
	if err != nil {
		return nil, err
	}
	autoservicioPOS, err := s.ParametroRepo.Consultar("AUTOSERVICIO_POS")
	if err != nil {
		return nil, err
	}

	// 2. Lógica condicional
	if autoservicioPOS != nil && autoservicioPOS.Valor == "S" {
		if autoservicioMaestro != nil && autoservicioMaestro.Valor == "S" {
			// SI “AUTOSERVICIO_MAESTRO” es “S” hacer una petición a la base de datos
			// Llamada a la función almacenada fnc_obtener_configuracion_pos_autoservicio
			query := "SELECT * FROM public.fnc_obtener_configuracion_pos_autoservicio($1)"
			// Se pasa un JSON vacío como parámetro ya que la función almacenada no lo utiliza actualmente
			args := []interface{}{"{}"}

			result, err := s.ParametroRepo.Client.Select(query, args)
			if err != nil {
				log.Printf("Error al ejecutar la función almacenada: %v", err)
				return nil, errors.New("Error al obtener la configuración desde la base de datos")
			}

			if len(result) == 0 || len(result[0]) == 0 {
				return nil, errors.New("No se obtuvo respuesta de la función almacenada")
			}

			jsonResponseStr, ok := result[0][0].(string)
			if !ok {
				log.Printf("Tipo de dato inesperado de la función almacenada: %T", result[0][0])
				return nil, errors.New("Formato de respuesta inesperado de la función almacenada")
			}

			var configData comunes_entidades.ConfiguracionInicial // Asumiendo que ConfiguracionInicial puede mapear la respuesta
			if err := json.Unmarshal([]byte(jsonResponseStr), &configData); err != nil {
				log.Printf("Error al parsear la respuesta JSON de la función almacenada: %v", err)
				return nil, errors.New("Error al parsear la respuesta JSON de la función almacenada")
			}
			return &configData, nil

		} else if autoservicioMaestro != nil && autoservicioMaestro.Valor == "N" {
			// SI “AUTOSERVICIO_MAESTRO” es “N” debe enviar una petición al POS MAESTRO
			if autoservicioIPMaestro == nil || autoservicioIPMaestro.Valor == "" {
				return nil, errors.New("AUTOSERVICIO_IP_MAESTRO no configurado para petición a POS MAESTRO")
			}
			url := autoservicioIPMaestro.Valor + ":" + constants.HOST_PORT + constants.API_PATH + constants.API_CONFIGURACION + constants.API_POS_MAESTRO
			_, err := s.HTTPClient.Send(
				"GET",
				url,
				&entities.HttpRequest{},
			)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Valor de AUTOSERVICIO_MAESTRO no válido o no encontrado")
		}
	} else if autoservicioPOS != nil && autoservicioPOS.Valor == "N" {
		// SI “AUTOSERVICIO_POS” es “N” No debe ejecutar ninguna acción y devolver un JSON
		response := map[string]string{"mensaje": "El POS no es autoservicio"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		// Aquí se debería devolver el JSON como parte de la respuesta HTTP, no directamente desde el servicio.
		// Para propósitos de demostración, se devuelve un error con el mensaje.
		return nil, errors.New(string(jsonResponse))
	} else {
		return nil, errors.New("Valor de AUTOSERVICIO_POS no válido o no encontrado")
	}

	return s.UseCase.Execute()
}
