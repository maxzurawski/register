package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/register/config"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/publishing"
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
