package persistence

import (
	"golang-queues/pkg/config/properties"
	"log"
	"sync"
)

var (
	instance *MQImpl
	once     sync.Once
)

type MQImpl struct {
	prop properties.DataBaseInfo
}

func GetMQImplInstance(p properties.DataBaseInfo) *MQImpl {
	once.Do(func() {
		instance = &MQImpl{p}
	})
	return instance
}

func (t *MQImpl) Save(msg string) error {
	log.Printf("saving message '%s' in postgres database named '%s'.\n", msg, t.prop.Name)
	return nil
}
