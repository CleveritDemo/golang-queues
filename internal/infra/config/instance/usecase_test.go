package instance

import (
	"github.com/gregperez/loadinitms"
	"reflect"
	"testing"
)

func TestGetRabbitMQUseCase(t *testing.T) {
	tests := []struct {
		name   string
		noWant loadinitms.PrimaryProcess
	}{
		{
			"TestGetRabbitMQUseCase",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRabbitMQUseCase(); reflect.DeepEqual(got, tt.noWant) {
				t.Errorf("GetRabbitMQUseCase() = %v, want %v", got, tt.noWant)
			}
		})
	}
}
