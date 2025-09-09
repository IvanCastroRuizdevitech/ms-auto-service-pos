package domain_repositories_obtener_parametros

import comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"

type IObtenerParametrizacion interface {
	Consultar(codigo string) (*comunes_entidades.ParametrosWatcher, error)
}
