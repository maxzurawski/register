package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/register/config"
	"github.com/maxzurawski/utilities/rabbit/crosscutting"
	"github.com/maxzurawski/utilities/rabbit/publishing"
)

type sensorsPublisher struct {
	*publishing.Publisher
}

func (p *sensorsPublisher) PublishDeleteChange(previous, current interface{}) {
	p.publishRaw("delete", previous, current)
}

func (p *sensorsPublisher) PublishSaveChange(previous, current interface{}) {
	p.publishRaw("save", previous, current)
}

func (p *sensorsPublisher) PublishUpdateChange(previous, current interface{}) {
	p.publishRaw("update", previous, current)
}

func (p *sensorsPublisher) publishRaw(routingKeySuffix string, previous, current interface{}) {
	if !config.Config().ConnectToRabbit() {
		log.Info("connection to rabbit disabled")
		return
	}
	p.Reset()
	p.PublishConfigurationChanged(crosscutting.RoutingKeySensors.String()+"."+routingKeySuffix,
		config.Config().ServiceName(),
		previous,
		current,
	)
}
