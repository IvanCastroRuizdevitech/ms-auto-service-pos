package aplication_services_configuracion_esclavo

import (
	aplication_usecases_configuracion_esclavo "genexis/pos/autoservicios/aplication/usecases/configuracion_esclavo"
	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	infrastructura_repositorios "genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	"log"
)

type GetConfiguracionInicialEsclavoService struct {
	UseCase       *aplication_usecases_configuracion_esclavo.GetConfiguracionInicialEsclavoUseCase
	HTTPClient    domain_adapters_clients_http.IClientHttp
	ParametroRepo *infrastructura_repositorios.RecuperarParametrosPos
}

func (s *GetConfiguracionInicialEsclavoService) Execute(body *comunes_entidades.ConfiguracionInicialEsclavo) (*comunes_entidades.ConfiguracionInicial, error) {
	log.Print("CONSULTANDO CONFIGURACION INICIAL POS ESCLAVO: ", body)
	return s.UseCase.Execute(body)
}
