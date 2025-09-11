package aplication_services_configuracion

import (
	"encoding/json"
	"errors"
	aplication_usecases_configuracion "genexis/pos/autoservicios/aplication/usecases/configuracion"
	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
	"genexis/pos/autoservicios/domain/constants"
	"genexis/pos/autoservicios/domain/entities"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	infrastructura_repositorios "genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	"log"
)

type GetConfiguracionInicialService struct {
	UseCase       *aplication_usecases_configuracion.GetConfiguracionInicialUseCase
	HTTPClient    domain_adapters_clients_http.IClientHttp
	ParametroRepo *infrastructura_repositorios.RecuperarParametrosPos
}

func (s *GetConfiguracionInicialService) Execute() (*comunes_entidades.ConfiguracionInicial, error) {

	log.Printf("CONSULTANDO CONFIGURACION INICIAL")
	autoservicioMaestro, err := s.ParametroRepo.Consultar(constants.AUTOSERVICIO_MAESTRO)
	if err != nil {
		return nil, err
	}
	autoservicioIPMaestro, err := s.ParametroRepo.Consultar(constants.AUTOSERVICIO_IP_MAESTRO)
	if err != nil {
		return nil, err
	}
	autoservicioPOS, err := s.ParametroRepo.Consultar(constants.AUTOSERVICIO_POS)
	if err != nil {
		return nil, err
	}
	autoservicioCARAS, err := s.ParametroRepo.Consultar(constants.AUTOSERVICIO_CARAS)
	if err != nil {
		return nil, err
	}

	if autoservicioPOS != nil && autoservicioPOS.Valor == "S" {
		if autoservicioMaestro != nil && autoservicioMaestro.Valor == "S" {
			caras := &comunes_entidades.ConfiguracionInicialEsclavo{
				Caras: autoservicioCARAS.Valor,
			}
			log.Print("caras: ", caras)
			return s.UseCase.Execute(caras)
		} else if autoservicioMaestro != nil && autoservicioMaestro.Valor == "N" {
			if autoservicioIPMaestro == nil || autoservicioIPMaestro.Valor == "" {
				return nil, errors.New("AUTOSERVICIO_IP_MAESTRO no configurado para petición a POS MAESTRO")
			}
			//url := autoservicioIPMaestro.Valor + ":" + constants.HOST_PORT + constants.API_PATH + constants.API_CONFIGURACION + constants.API_POS_MAESTRO
			url := "localhost" + ":" + constants.HOST_PORT + constants.API_PATH + constants.API_CONFIGURACION + constants.API_POS_MAESTRO
			bodyMarshal, _ := json.Marshal(comunes_entidades.ConfiguracionInicialEsclavo{Caras: autoservicioCARAS.Valor})
			result, err := s.HTTPClient.Send(
				"POST",
				url,
				&entities.HttpRequest{
					Body: bodyMarshal,
				},
			)
			if err != nil {
				return nil, err
			}

			var response *comunes_entidades.ConfiguracionInicial
			err = json.Unmarshal(result.Body, &response)
			if err != nil {
				return nil, err
			}
			return response, err
		} else {
			return nil, errors.New("Valor de AUTOSERVICIO_MAESTRO no válido o no encontrado")
		}
	} else if autoservicioPOS != nil && autoservicioPOS.Valor == "N" {
		response := map[string]string{"mensaje": "El POS no es autoservicio"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(jsonResponse))
	} else {
		return nil, errors.New("Valor de AUTOSERVICIO_POS no válido o no encontrado")
	}

	//return s.UseCase.Execute()
}
