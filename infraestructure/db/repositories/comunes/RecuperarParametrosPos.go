package infrastructura_repositorios

import (
	domain_adapters_clients_db "genexis/pos/autoservicios/domain/adapters/clients/db"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	"log"
	"strings"
)

type RecuperarParametrosPos struct {
	Client domain_adapters_clients_db.IClientDB
}

func (RWP *RecuperarParametrosPos) Consultar(codigo string) (*comunes_entidades.ParametrosWatcher, error) {
	args := []any{}
	query := `SELECT x.* FROM public.wacher_parametros x WHERE codigo ilike $1`
	query = strings.ReplaceAll(query, "$1", "'%"+codigo+"%'")
	log.Printf("CONSULTA PARAMETROS %s", query)
	respuesta, err := RWP.Client.Select(query, args)

	if err != nil {
		return nil, err
	}

	parametro := &comunes_entidades.ParametrosWatcher{}
	for _, valor := range respuesta {
		parametro.Id = valor[0].(int64)
		parametro.Codigo = valor[1].(string)
		parametro.Tipo = valor[2].(int32)
		parametro.Valor = valor[3].(string)
	}

	return parametro, nil
}
