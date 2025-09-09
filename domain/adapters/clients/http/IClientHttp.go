package dominio_adapters_clients_http

import "genexis/pos/autoservicios/domain/entities"

type IClientHttp interface {
	Send(method string, url string, message *entities.HttpRequest) (*entities.HttpResponse, error)
}
