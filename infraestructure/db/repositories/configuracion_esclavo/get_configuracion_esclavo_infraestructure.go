package infraestructura_repos_configuracion_esclavo

import (
	"encoding/json"
	domain_adapters_clients_db "genexis/pos/autoservicios/domain/adapters/clients/db"
	comunes_entidades "genexis/pos/autoservicios/domain/entities/entidades_comunes"
	"log"
)

type ConfiguracionInicialEsclavoRepository struct {
	Client domain_adapters_clients_db.IClientDB
}

func (r *ConfiguracionInicialEsclavoRepository) ConsultarConfiguracionEsclavo(body *comunes_entidades.ConfiguracionInicialEsclavo) (*comunes_entidades.ConfiguracionInicial, error) {
	log.Printf("CONSULTANDO CONFIGURACION INICIAL POS ESCLAVO")
	query := "SELECT * FROM public.fnc_obtener_configuracion_pos_autoservicio($1::jsonb);"
	results, err := r.Client.Exec(query, []any{body})
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil // O error si esperas siempre un resultado
	}

	var configuracion *comunes_entidades.ConfiguracionInicial

	marshal, err := json.Marshal(results[0][0])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(marshal, &configuracion)
	if err != nil {
		return nil, err
	}

	return configuracion, nil
}
