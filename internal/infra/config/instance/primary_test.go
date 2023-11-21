package instance

import (
	"github.com/gregperez/loadinitms"
	"reflect"
	"testing"
)

func TestGetRabbitMQSubscriberInstance(t *testing.T) {
	tests := []struct {
		name string
		want loadinitms.PrimaryProcess
	}{
		{
			name: "Test GetRabbitMQSubscriberInstance",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRabbitMQSubscriberInstance(); reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubDSOrderStatusInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
