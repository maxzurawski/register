package publishers

import (
	"github.com/labstack/gommon/log"
	"github.com/xdevices/register/config"
	"github.com/xdevices/utilities/rabbit/crosscutting"
	"github.com/xdevices/utilities/rabbit/publishing"
)

type sensorsPublisher struct {
	intern *publishing.Publisher
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
	p.intern.PublishConfigurationChanged(crosscutting.RoutingKeySensors.String()+"."+routingKeySuffix,
		config.Config().ServiceName(),
		previous,
		current,
	)
}
