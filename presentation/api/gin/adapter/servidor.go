package api_adaptador_servidor

import (
	"genexis/pos/autoservicios/domain/constants"
	api_routes "genexis/pos/autoservicios/presentation/api/gin/routes"
	"log"
)

func Start() error {

	servidor, err := api_routes.GinConfig()

	if err != nil {
		log.Fatal(err)
		return err
	}

	ip := constants.HOST_IP
	port := constants.HOST_PORT
	host := ip + ":" + port
	err = servidor.Run(host)
	if err != nil {
		log.Fatal(err)
		return err

	}
	log.Println("SERVIDOR CORRIENDO: ", host)

	return nil
}
