package aplication_usecases_configuracion

import (
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	domain_repositories_configuracion "genexis/pos/autoservicios/domain/repositories/db/configuracion"
)

type GetConfiguracionInicialUseCase struct {
	Repository domain_repositories_configuracion.IConfiguracionInicial
}

func (uc *GetConfiguracionInicialUseCase) Execute() (*comunes_entidades.ConfiguracionInicial, error) {
	return uc.Repository.ConsultarConfiguracion()
}
