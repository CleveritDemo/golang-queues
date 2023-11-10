package core

import "sync"

var (
	instance *RabbitMQImpl
	once     sync.Once
)

type RabbitMQImpl struct {
	persistencePort PersistenceMQPort
}

func GetRabbitMQImpl(persistencePort PersistenceMQPort) *RabbitMQImpl {
	once.Do(func() {
		instance = &RabbitMQImpl{
			persistencePort: persistencePort,
		}
	})
	return instance
}

func (t *RabbitMQImpl) Handle(msg string) error {
	return t.persistencePort.Save(msg)
}
