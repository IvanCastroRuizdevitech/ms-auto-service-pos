package aplication_services_configuracion

import (
	aplication_usecases_configuracion "genexis/pos/autoservicios/aplication/usecases/configuracion"
	"genexis/pos/autoservicios/domain/entities"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
	infrastructura_repositorios "genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	"encoding/json"
	"errors"
)

type GetConfiguracionInicialService struct {
	UseCase    *aplication_usecases_configuracion.GetConfiguracionInicialUseCase
	HTTPClient domain_adapters_clients_http.IClientHttp
	ParametroRepo *infrastructura_repositorios.RecuperarParametrosPos
}

func (s *GetConfiguracionInicialService) Execute() (*comunes_entidades.ConfiguracionInicial, error) {

	// 1. Recuperar parámetros
	autoservicioMaestro, err := s.ParametroRepo.Consultar("AUTOSERVICIO_MAESTRO")
	if err != nil { return nil, err }
	autoservicioIPMaestro, err := s.ParametroRepo.Consultar("AUTOSERVICIO_IP_MAESTRO")
	if err != nil { return nil, err }
	autoservicioPOS, err := s.ParametroRepo.Consultar("AUTOSERVICIO_POS")
	if err != nil { return nil, err }

	// 2. Lógica condicional
	if autoservicioPOS != nil && autoservicioPOS.Valor == "S" {
		if autoservicioMaestro != nil && autoservicioMaestro.Valor == "S" {
			// SI “AUTOSERVICIO_MAESTRO” es “S” hacer una petición a la base de datos
			// TODO: Implementar la llamada a la función almacenada para obtener datos de surtidores, productos, etc.
			// Por ahora, se devuelve un error o una respuesta de ejemplo.
			return nil, errors.New("Funcionalidad de base de datos para AUTOSERVICIO_MAESTRO = S pendiente de implementar")
		} else if autoservicioMaestro != nil && autoservicioMaestro.Valor == "N" {
			// SI “AUTOSERVICIO_MAESTRO” es “N” debe enviar una petición al POS MAESTRO
			if autoservicioIPMaestro == nil || autoservicioIPMaestro.Valor == "" {
				return nil, errors.New("AUTOSERVICIO_IP_MAESTRO no configurado para petición a POS MAESTRO")
			}
			url := autoservicioIPMaestro.Valor + "/configuracion"
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


