package aplication_usecases_configuracion_esclavo

import (
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	domain_repositories_configuracion_esclavo "genexis/pos/autoservicios/domain/repositories/db/configuracion_esclavo"
)

type GetConfiguracionInicialEsclavoUseCase struct {
	Repository domain_repositories_configuracion_esclavo.IConfiguracionInicialEsclavo
}

func (uc *GetConfiguracionInicialEsclavoUseCase) Execute() (*comunes_entidades.ConfiguracionInicial, error) {
	return uc.Repository.ConsultarConfiguracionEsclavo()
}
