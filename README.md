# ms-auto-service-pos
Servicio encargado de Gestionar los procesos del POS de Autoservicios.

# Para renderizado automatico en dev
- go install github.com/air-verse/air@latest
- air init
- Iniciar el proyecto con : air

# Compilar
- GOOS=linux GOARCH=arm64 go build -o ms-auto-service-pos main.go
