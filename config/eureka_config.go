package config

import "github.com/maxzurawski/utilities/discovery"

type EurekaManager struct {
	discovery.Manager
}

func EurekaManagerInit() *EurekaManager {
	manager := EurekaManager{
		Manager: discovery.Manager{
			RegistrationTicket: Config().RegistrationTicket(),
			EurekaService:      Config().EurekaService(),
		},
	}
	return &manager
}
