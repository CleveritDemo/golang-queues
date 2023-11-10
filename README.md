# Arquetipo Golang con Hexagonal y Consumo de Colas RabbitMQ

Este es un arquetipo básico de una aplicación en Go que sigue la arquitectura hexagonal y utiliza RabbitMQ para el consumo de colas.

## Estructura del Proyecto

La estructura del proyecto sigue una arquitectura hexagonal para mantener una clara separación entre las capas de la aplicación.

```plaintext
golang-queues/
  ├── cmd/
  │   ├──  main.go
  ├── internal/
  │   ├── core/
  │   │   ├── port.go
  │   │   ├── service.go
  |   |   └── usecase.go
  │   │   └── ...
  │   ├── infra/
  │   │   ├── config/
  │   │   │   ├── instance/
  │   │   │   │   ├── primary.go
  │   │   │   │   └── ...
  │   │   ├── primary/
  │   │   │   ├── subscriber.go
  │   │   │   └── ...
  │   │   ├── secondary/
  │   │   │   ├── publisher.go
  │   │   │   └── ...
  │   ├── resource/
  │       ├── properties.yaml
  │       └── ...
  ├── pkg/
  ├── go.mod
  └── go.sum
```

## Configuración e Inicialización

Asegúrate de tener RabbitMQ instalado y en ejecución en tu máquina o en un clúster remoto antes de ejecutar la aplicación.

```bash
# Instalar las dependencias del proyecto
go mod tidy

# Ejecutar rabbitmq en docker
docker run -it --rm --name rabbit-test -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin
```

## Ejecución

```bash
# Ejecutar la aplicación
go run cmd/main.go
```

## Contribuciones

¡Las contribuciones son bienvenidas! Si encuentras algún problema o tienes mejoras para proponer, por favor, crea un issue o envía un pull request.