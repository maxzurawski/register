package sensor

import "github.com/maxzurawski/register/dto"

func (s *service) GetCachedSensors() ([]dto.CachedSensorDTO, error) {
	sensors, err := s.mgr.GetAllSensors()

	if err != nil {
		return nil, err
	}

	var cdtos []dto.CachedSensorDTO
	for _, item := range sensors {
		sensorDTO := s.mgr.MapToCachedSensorDTO(&item)
		cdtos = append(cdtos, sensorDTO)
	}
	return cdtos, nil
}
