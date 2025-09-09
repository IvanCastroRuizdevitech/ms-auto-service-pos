package aplication_services_configuracion

import (
	aplication_usecases_configuracion "genexis/pos/autoservicios/aplication/usecases/configuracion"
	"genexis/pos/autoservicios/domain/entities"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
)

type GetConfiguracionInicialService struct {
	UseCase    *aplication_usecases_configuracion.GetConfiguracionInicialUseCase
	HTTPClient domain_adapters_clients_http.IClientHttp
}

func (s *GetConfiguracionInicialService) Execute() (*comunes_entidades.ConfiguracionInicial, error) {

	_, err := s.HTTPClient.Send(
		"GET",
		"https://api.ejemplo.com/configuracion",
		&entities.HttpRequest{},
	)
	if err != nil {
		return nil, err
	}

	return s.UseCase.Execute()
}
