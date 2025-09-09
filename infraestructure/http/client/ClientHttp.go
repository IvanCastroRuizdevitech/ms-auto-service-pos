package cliente_infrastruc

import (
	"bytes"
	"genexis/pos/autoservicios/domain/entities"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type ClientHttp struct {
	client *http.Client
}

func (CHttp *ClientHttp) Send(method string, url string, message *entities.HttpRequest) (*entities.HttpResponse, error) {

	request, err := http.NewRequest(method, "http://"+url, bytes.NewBuffer(message.Body))
	if err != nil {
		log.Printf("ERROR DEL HOS %s : %v", url, err)
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := CHttp.client.Do(request)

	if err != nil {
		log.Printf("DESPUES DEL DO %v \n", err)
		errorNet := err.(net.Error)
		return nil, errorNet
	}

	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	responesdomain := &entities.HttpResponse{
		StatusCode: response.StatusCode,
		Body:       bodyBytes,
		Status:     response.Status,
	}

	log.Printf("RESPUESTA DE PETICION %s : %d \n", url, responesdomain.StatusCode)
	log.Printf("BODY : %v \n", string(bodyBytes))
	return responesdomain, nil
}

func InitializeClient() (*ClientHttp, error) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	return &ClientHttp{
		client: httpClient,
	}, nil
}
