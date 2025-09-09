package entities

import (
	"log"
	"net/http"
)

type ServerResponse[T any] struct {
	Success bool   `json:"success,omitempty"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func NewSuccessServerResponse(message string, data interface{}) ServerResponse[any] {
	return ServerResponse[any]{
		Success: true,
		Message: message,
		Data:    data,
	}
}
func NewServerResponseDataBase(message string, data interface{}, success bool) ServerResponse[any] {
	return ServerResponse[any]{
		Success: success,
		Message: message,
		Data:    data,
	}
}

func NewErrorServerResponse(message string, err error) ServerResponse[any] {
	log.Println("[HTTP LOCAL] NEW ERROR RESPONSE...", err.Error())
	return ServerResponse[any]{
		Status:  http.StatusBadRequest,
		Success: false,
		Message: message,
	}
}
