package core

type PersistenceMQPort interface {
	Save(msg string) error
}
