package domain_repositories_configuracion

import comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"

type IConfiguracionInicial interface {
	ConsultarConfiguracion(body *comunes_entidades.ConfiguracionInicialEsclavo) (*comunes_entidades.ConfiguracionInicial, error)
}
