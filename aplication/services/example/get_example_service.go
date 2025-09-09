package aplication_services_example

import (
	entities_example "genexis/pos/autoservicios/domain/entities/example"
	domain_usecases_example "genexis/pos/autoservicios/domain/usecases/example"
)

type GetExampleService struct {
	GetExampleUseCase domain_usecases_example.GetExampleUseCase
}

func (GES *GetExampleService) Execute(param int) (*[]entities_example.DataResponse, error) {
	return GES.GetExampleUseCase.Execute(param)
}
