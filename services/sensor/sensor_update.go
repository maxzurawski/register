package sensor

import (
	"github.com/xdevices/register/dto"
)

func (s *service) Update(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error) {
	updated, err := s.mgr.UpdateSensor(registerDTO)
	dto := s.mgr.MapToSensorDTO(updated)
	return &dto, err
}
