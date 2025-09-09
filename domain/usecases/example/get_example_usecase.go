package domain_usecases_example

import (
	entities_example "genexis/pos/autoservicios/domain/entities/example"
)

type GetExampleUseCase interface {
	Execute(param int) (*[]entities_example.DataResponse, error)
}
