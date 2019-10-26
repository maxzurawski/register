package publishers

import (
	"github.com/maxzurawski/register/config"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/publishing"
)

var configChangedPublisher *publishing.Publisher
var sensorsPublisherInstance *sensorsPublisher
var attributesPublisherInstance *attributesPublisher

func Init() {
	if configChangedPublisher == nil && config.Config().ConnectToRabbit() {
		configChangedPublisher = config.Config().InitPublisher()
		// NOTE: once declared - even if we disconnect, exchange will stay there in rabbitmq
		configChangedPublisher.DeclareTopicExchange(crosscutting.TopicConfigurationChanged.String())
	}
}

func SensorsPublisher() *sensorsPublisher {
	if sensorsPublisherInstance == nil {
		sensorsPublisherInstance = &sensorsPublisher{
			configChangedPublisher,
		}
	}
	return sensorsPublisherInstance
}

func AttributesPublisher() *attributesPublisher {
	if attributesPublisherInstance == nil {
		attributesPublisherInstance = new(attributesPublisher)
		attributesPublisherInstance.Publisher = configChangedPublisher
	}
	return attributesPublisherInstance
}
