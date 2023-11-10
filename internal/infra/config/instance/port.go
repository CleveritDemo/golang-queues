package instance

import (
	"golang-queues/internal/core"
	"golang-queues/internal/infra/config/properties"
	"golang-queues/internal/infra/secondary/persistence"
)

func GetPersistencePort() core.PersistenceMQPort {
	return persistence.GetMQImplInstance(
		properties.GetDatabaseProp().Postgresql,
	)
}
