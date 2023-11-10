package subscriber

import (
	"golang-queues/internal/core"
	"golang-queues/pkg/config/properties"
	"golang-queues/pkg/rabbitmq"
	"log"
	"sync"
)

var (
	subSingleton *JohnDoeMQSubscription
	once         sync.Once
)

type JohnDoeMQSubscription struct {
	subProperties properties.RabbitMQInfo
	useCase       core.RabbitMQUseCase
}

func GetMqSubscriptionInstance(p properties.RabbitMQInfo, u core.RabbitMQUseCase) *JohnDoeMQSubscription {
	once.Do(func() {
		subSingleton = &JohnDoeMQSubscription{p, u}
	})
	return subSingleton
}

func (s *JohnDoeMQSubscription) Start() {
	log.Println("Starting John Doe subscription to RabbitMQ")
	rabbitmq.StartConsumer(s.subProperties.Queue, s, &s.subProperties)
}

// HandleSubEvent handles the event
func (s *JohnDoeMQSubscription) HandleSubEvent(event string) {
	err := s.useCase.Handle(event)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("event trigger success")
	}
}
