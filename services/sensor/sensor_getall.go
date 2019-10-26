package sensor

import "github.com/xdevices/register/dto"

func (s *service) GetAll() ([]dto.SensorRegisterDTO, error) {
	// NOTE: get all entities with dbmanager
	sensors, err := s.mgr.GetAllSensors()
	if err != nil {
		return nil, err
	}

	var dtos []dto.SensorRegisterDTO
	// NOTE: map entities into dtos
	for _, item := range sensors {
		dtos = append(dtos, s.mgr.MapToSensorDTO(&item))
	}
	return dtos, nil
}
