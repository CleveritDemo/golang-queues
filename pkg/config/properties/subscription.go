package properties

import "sync"

var (
	subscriptionPropsInstance     *SubscriptionProperty
	subscriptionPropsInstanceOnce sync.Once
)

func GetSubscriptionProperty() *SubscriptionProperty {
	subscriptionPropsInstanceOnce.Do(func() {
		subscriptionPropsInstance = &SubscriptionProperty{}
	})
	return subscriptionPropsInstance
}

// SubscriptionProperty Property Struct Definition
type SubscriptionProperty struct {
	Subscriptions `yaml:"subscriptions"`
}

type Subscriptions struct {
	JohnDoe RabbitMQInfo `yaml:"john-doe"`
}

type RabbitMQInfo struct {
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Exchange       string `yaml:"exchange"`
	Queue          string `yaml:"queue"`
	RoutingKey     string `yaml:"routing-key"`
	ConsumerTag    string `yaml:"consumer-tag"`
	WorkerPoolSize int    `yaml:"worker-pool-size"`
}
