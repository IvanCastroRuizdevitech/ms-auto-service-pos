package casosuso_comunes

import (
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	domain_repositories_obtener_parametros "genexis/pos/autoservicios/domain/repositories/db/wacher_parametros"
)

type ObtenerParametroWatcher struct {
	ObtenerParametroWatcher domain_repositories_obtener_parametros.IObtenerParametrizacion
}

func (OPW *ObtenerParametroWatcher) Ejecutar(codigo string) (*comunes_entidades.ParametrosWatcher, error) {
	return OPW.ObtenerParametroWatcher.Consultar(codigo)
}
