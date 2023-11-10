package instance

import (
	"golang-queues/internal/core"
)

func GetRabbitMQUseCase() core.RabbitMQUseCase {
	return core.GetRabbitMQImpl(
		GetPersistencePort(),
	)
}
