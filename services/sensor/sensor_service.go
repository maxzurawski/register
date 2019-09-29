package sensor

import (
	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/dto"
)

// here does not mean we have to ask in each method database
type SensorsService interface {
	Save(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error)
}

var Service SensorsService

type service struct {
	mgr dbprovider.DBManager
}

func Init() {
	s := service{}
	s.mgr = dbprovider.Mgr
	Service = &s
}
