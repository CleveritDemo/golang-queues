package core

import (
	"errors"
	"reflect"
	"testing"
)

type MockPersistenceMQPort struct {
	wantErr bool
}

func (m MockPersistenceMQPort) Save(msg string) error {
	if m.wantErr {
		return errors.New("error in Save")
	}
	return nil
}

func TestGetRabbitMQImpl_Instance(t *testing.T) {
	type args struct {
		save PersistenceMQPort
	}
	tests := []struct {
		name string
		args args
		want RabbitMQImpl
	}{
		{
			name: "test GetRabbitMQImpl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRabbitMQImpl(tt.args.save); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRabbitMQImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRabbitMQImpl_Handle(t *testing.T) {
	type fields struct {
		persistencePort PersistenceMQPort
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test RabbitMQImpl",
			fields{MockPersistenceMQPort{}},
			args{"test"},
			false,
		},
		{
			"test RabbitMQImpl with error",
			fields{MockPersistenceMQPort{true}},
			args{"test"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RabbitMQImpl{
				persistencePort: tt.fields.persistencePort,
			}
			if err := r.Handle(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("RabbitMQImpl.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
