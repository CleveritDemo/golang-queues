package main

import (
	"github.com/gregperez/loadinitms"
	"golang-queues/internal/infra/config/instance"
	"golang-queues/internal/infra/config/properties"
	properties2 "golang-queues/pkg/config/properties"
)

// main.go
func main() {

	// Setting properties file path
	loadinitms.SetPropertyFilePath("internal/resources/properties.yml")

	// Adding primaries
	loadinitms.AddPrimary(instance.GetRabbitMQSubscriberInstance)

	// Adding properties
	loadinitms.AddProperty(properties2.GetSubscriptionProperty())
	loadinitms.AddProperty(properties.GetDatabaseProp())

	// Running loadinitms
	loadinitms.Run()
}
