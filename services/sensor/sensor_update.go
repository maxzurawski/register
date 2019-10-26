package sensor

import (
	"github.com/maxzurawski/register/dto"
)

func (s *service) Update(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error) {
	updated, err := s.mgr.UpdateSensor(registerDTO)
	dto := s.mgr.MapToSensorDTO(updated)
	return &dto, err
}
