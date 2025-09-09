package aplication_usecases_example

import (
	entities_example "genexis/pos/autoservicios/domain/entities/example"
	domain_repositories_example "genexis/pos/autoservicios/domain/repositories/db/example"
)

type GetExampleUseCase struct {
	Repository domain_repositories_example.GetExampleInterface
}

func (GEUC *GetExampleUseCase) Execute(param int) (*[]entities_example.DataResponse, error) {
	return GEUC.Repository.Get(param)
}
