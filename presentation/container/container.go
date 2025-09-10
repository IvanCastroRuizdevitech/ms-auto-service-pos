package container

import (
	aplication_services_configuracion "genexis/pos/autoservicios/aplication/services/configuracion"
	aplication_services_configuracion_esclavo "genexis/pos/autoservicios/aplication/services/configuracion_esclavo"
	aplication_usecases_configuracion "genexis/pos/autoservicios/aplication/usecases/configuracion"
	aplication_usecases_configuracion_esclavo "genexis/pos/autoservicios/aplication/usecases/configuracion_esclavo"
	infraestructura_repos_comunes "genexis/pos/autoservicios/infraestructure/db/repositories/comunes"
	infraestructura_repos_configuracion "genexis/pos/autoservicios/infraestructure/db/repositories/configuracion"
	infraestructura_repos_configuracion_esclavo "genexis/pos/autoservicios/infraestructure/db/repositories/configuracion_esclavo"

	aplication_services_example "genexis/pos/autoservicios/aplication/services/example"
	"genexis/pos/autoservicios/aplication/usecases/casosuso_comunes"

	aplication_usecases_example "genexis/pos/autoservicios/aplication/usecases/example"

	domain_adapters_clients_http "genexis/pos/autoservicios/domain/adapters/clients/http"
	"genexis/pos/autoservicios/domain/constants"

	domain_repositories_example "genexis/pos/autoservicios/domain/repositories/db/example"
	domain_repositories_obtener_parametros "genexis/pos/autoservicios/domain/repositories/db/wacher_parametros"

	domain_usecases_example "genexis/pos/autoservicios/domain/usecases/example"

	infraestructure_db_client "genexis/pos/autoservicios/infraestructure/db/client"
	infraestructure_repos_example "genexis/pos/autoservicios/infraestructure/db/repositories/example"

	infraestructure_http_client "genexis/pos/autoservicios/infraestructure/http/client"

	"log"
)

// SERVICES
var GetExampleService *aplication_services_example.GetExampleService
var GetConfiguracionInicialService *aplication_services_configuracion.GetConfiguracionInicialService
var GetConfiguracionInicialEsclavoService *aplication_services_configuracion_esclavo.GetConfiguracionInicialEsclavoService

// USE CASES
var GetExampleUseCase domain_usecases_example.GetExampleUseCase
var ObtenerWatcherParametros *casosuso_comunes.ObtenerParametroWatcher

// REPOSITORIES
var GetExampleRepository domain_repositories_example.GetExampleInterface
var RecuperarWatcherParametrosRepo domain_repositories_obtener_parametros.IObtenerParametrizacion

// GENERAL

// CLIENTS

var clientHttp domain_adapters_clients_http.IClientHttp

// MAPPER

func InitializeContainer() error {

	//MAPPER

	//clients

	clientDB, err := infraestructure_db_client.InitializeClient(constants.DB_CON)
	if err != nil {
		log.Fatal("Error inicializar bd client", err)
		return err

	}

	clientHttp, err = infraestructure_http_client.InitializeClient()

	if err != nil {
		log.Fatal(err)
		return err

	}

	// REPOSITORIES
	GetExampleRepository = &infraestructure_repos_example.GetExampleRepository{
		Client: clientDB,
	}
	GetConfiguracionRepository := &infraestructura_repos_configuracion.ConfiguracionInicialRepository{
		Client: clientDB,
	}
  RecuperarWatcherParametrosRepo := &infraestructura_repos_comunes.RecuperarParametrosPos{
		Client: clientDB,
	}
	GetConfiguracionEsclavoRepository := &infraestructura_repos_configuracion_esclavo.ConfiguracionInicialEsclavoRepository{
		Client: clientDB,
	}

	// USE CASES
	GetExampleUseCase = &aplication_usecases_example.GetExampleUseCase{
		Repository: GetExampleRepository,
	}
	GetConfiguracionUseCase := &aplication_usecases_configuracion.GetConfiguracionInicialUseCase{
		Repository: GetConfiguracionRepository,
	}
	GetConfiguracionEsclavoUseCase := &aplication_usecases_configuracion_esclavo.GetConfiguracionInicialEsclavoUseCase{
		Repository: GetConfiguracionEsclavoRepository,
	}

	// SERVICES
	GetExampleService = &aplication_services_example.GetExampleService{
		GetExampleUseCase: GetExampleUseCase,
	}
	GetConfiguracionInicialService = &aplication_services_configuracion.GetConfiguracionInicialService{
		UseCase:       GetConfiguracionUseCase,
		HTTPClient:    clientHttp,
		ParametroRepo: RecuperarWatcherParametrosRepo,
	}
	GetConfiguracionInicialEsclavoService = &aplication_services_configuracion_esclavo.GetConfiguracionInicialEsclavoService{
		UseCase:    GetConfiguracionEsclavoUseCase,
		HTTPClient: clientHttp,
	}

	return nil
}
