package sensor

import (
	"github.com/xdevices/register/dto"
)

func (s *service) Save(registerDTO dto.SensorRegisterDTO) (*dto.SensorRegisterDTO, error) {
	register, err := s.mgr.SaveSensor(registerDTO)
	dto := s.mgr.MapToSensorDTO(register)
	return &dto, err
}
