package sensor

import (
	"github.com/maxzurawski/register/dbprovider"
	"github.com/maxzurawski/register/dto"
)

// here does not mean we have to ask in each method database
type SensorsService interface {
	Save(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error)
	Update(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error)
	Delete(uuid string) (uint, error)
	GetAll() ([]dto.SensorRegisterDTO, error)
	FindSensorByUuid(uuid string) *dto.SensorRegisterDTO

	// caches
	GetCachedSensors() ([]dto.CachedSensorDTO, error)
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
