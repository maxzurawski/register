package sensor

import (
	"github.com/labstack/gommon/log"
	"github.com/maxzurawski/register/dto"
)

func (s *service) FindSensorByUuid(uuid string) *dto.SensorRegisterDTO {

	// NOTE: search for specific sensor entity by uuid
	sensorRegister, err := s.mgr.GetSensorByUuid(uuid)
	if err != nil {
		log.Error(err)
		return nil
	}

	// NOTE: map entity into dto
	dto := s.mgr.MapToSensorDTO(sensorRegister)
	return &dto
}
