package core

type RabbitMQUseCase interface {
	Handle(msg string) error
}
