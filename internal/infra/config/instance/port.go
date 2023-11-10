package instance

import (
	"golang-queues/internal/core"
	"golang-queues/internal/infra/secondary/persistence"
	"golang-queues/pkg/config/properties"
)

func GetPersistencePort() core.PersistenceMQPort {
	return persistence.GetMQImplInstance(
		properties.GetDatabaseProp().Postgresql,
	)
}
