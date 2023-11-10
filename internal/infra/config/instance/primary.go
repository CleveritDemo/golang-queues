package instance

import (
	"github.com/gregperez/loadinitms"
	"golang-queues/internal/infra/primary/subscriber"
	"golang-queues/pkg/config/properties"
)

func GetRabbitMQSubscriberInstance() loadinitms.PrimaryProcess {
	return subscriber.GetMqSubscriptionInstance(
		properties.GetSubscriptionProperty().Subscriptions.JohnDoe,
		GetRabbitMQUseCase(),
	)
}
