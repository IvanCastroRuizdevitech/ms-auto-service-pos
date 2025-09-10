package domain_repositories_configuracion_esclavo

import comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"

type IConfiguracionInicialEsclavo interface {
	ConsultarConfiguracionEsclavo() (*comunes_entidades.ConfiguracionInicial, error)
}
