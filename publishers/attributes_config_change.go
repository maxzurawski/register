package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/xdevices/register/config"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/publishing"
)

type attributesPublisher struct {
	*publishing.Publisher
}

func (p *attributesPublisher) PublishUpdateChange(previous, current interface{}) {

	if !config.Config().ConnectToRabbit() {
		log.Info("connection to rabbit disabled")
		return
	}
	p.Reset()
	p.PublishConfigurationChanged(crosscutting.RoutingKeySensors.String()+".update",
		config.Config().ServiceName(),
		previous,
		current,
	)
}
