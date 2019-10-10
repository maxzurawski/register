package publishers

import (
	"github.com/xdevices/register/config"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/publishing"
)

var configChanged *publishing.Publisher
var sensorsPublisherInstance *sensorsPublisher

func Init() {
	if configChanged == nil && config.Config().ConnectToRabbit() {
		configChanged = config.Config().InitPublisher()
		// NOTE: once declared - even if we disconnect, exchange will stay there in rabbitmq
		configChanged.DeclareTopicExchange(crosscutting.TopicConfigurationChanged.String())
	}
}

func SensorsPublisher() *sensorsPublisher {
	if sensorsPublisherInstance == nil {
		sensorsPublisherInstance = new(sensorsPublisher)
		sensorsPublisherInstance.intern = configChanged
	}
	return sensorsPublisherInstance
}
