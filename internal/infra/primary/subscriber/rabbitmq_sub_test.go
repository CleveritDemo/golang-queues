package subscriber

import (
	"errors"
	"golang-queues/internal/core"
	"golang-queues/pkg/config/properties"
	"testing"
)

type MockSaveUseCase struct {
	wantErr bool
}

func (m MockSaveUseCase) Handle(msg string) error {
	if m.wantErr {
		return errors.New("error in Handle")
	}
	return nil
}

func TestGetMqSubscriptionInstance(t *testing.T) {
	mockUseCase := MockSaveUseCase{}
	propSub := properties.RabbitMQInfo{
		Queue: "test-queue",
	}

	subscriber := GetMqSubscriptionInstance(propSub, mockUseCase)

	if subscriber == nil {
		t.Error("Expect GetMqSubscriptionInstance not to be nil but got:", subscriber)
	}
}

func TestJohnDoeMQSubscription_HandleSubEvent(t *testing.T) {

	type fields struct {
		subProperties properties.RabbitMQInfo
		useCase       core.RabbitMQUseCase
	}

	tests := []struct {
		name   string
		fields fields
		msg    string
	}{
		{
			name: "Test HandleSubEvent successfully",
			fields: fields{
				subProperties: properties.RabbitMQInfo{
					Queue: "test-queue",
				},
				useCase: MockSaveUseCase{},
			},
			msg: "test-message",
		}, {
			name: "Test HandleSubEvent error in useCase",
			fields: fields{
				subProperties: properties.RabbitMQInfo{
					Queue: "test-queue",
				},
				useCase: MockSaveUseCase{true},
			},
			msg: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &JohnDoeMQSubscription{
				subProperties: tt.fields.subProperties,
				useCase:       tt.fields.useCase,
			}
			s.HandleSubEvent(tt.msg)
		})
	}
}
