package infraestructura_repos_example

import (
	"encoding/json"
	domain_adapters_clients_db "genexis/pos/autoservicios/domain/adapters/clients/db"
	entities_example "genexis/pos/autoservicios/domain/entities/example"
	"log"
)

type GetExampleRepository struct {
	Client domain_adapters_clients_db.IClientDB
}

func (GBPER *GetExampleRepository) Get(param int) (*[]entities_example.DataResponse, error) {

	query := "SELECT * FROM scheme.function($1);"
	results, err := GBPER.Client.Exec(query, []any{param})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var resutlsGeneral *[]entities_example.DataResponse
	jsonResults, err := json.Marshal(results[0][0])
	if err != nil {
		log.Println("❌ Error al serializar resultados:", err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(jsonResults), &resutlsGeneral); err != nil {
		log.Println("❌ Error al deserializar JSON general:", err)
		return nil, err
	}

	return resutlsGeneral, nil

}
