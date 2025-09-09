package domain_repositories_example

import entities_example "genexis/pos/autoservicios/domain/entities/example"

type GetExampleInterface interface {
	Get(param int) (*[]entities_example.DataResponse, error)
}
