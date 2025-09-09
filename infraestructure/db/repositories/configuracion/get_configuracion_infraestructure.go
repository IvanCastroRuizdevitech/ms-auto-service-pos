package infraestructura_repos_configuracion

import (
	domain_adapters_clients_db "genexis/pos/autoservicios/domain/adapters/clients/db"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
)

type ConfiguracionInicialRepository struct {
	Client domain_adapters_clients_db.IClientDB
}

func (r *ConfiguracionInicialRepository) ConsultarConfiguracion() (*comunes_entidades.ConfiguracionInicial, error) {
    query := "SELECT * FROM public.fnc_configuracion_inicial();"
    results, err := r.Client.Exec(query, []any{})
    if err != nil {
        return nil, err
    }

    if len(results) == 0 {
        return nil, nil // O error si esperas siempre un resultado
    }

    row := results[0]
    var configuracion comunes_entidades.ConfiguracionInicial

    // Asumiendo que los campos están en el mismo orden que la entidad
    // Ajusta los tipos según tu entidad
    configuracion.Id = row[0].(int64)         // Cambia el tipo si es necesario
    configuracion.Codigo = row[1].(string)    // Cambia el tipo si es necesario
    // ...agrega los demás campos según tu entidad y función almacenada

    return &configuracion, nil
}
