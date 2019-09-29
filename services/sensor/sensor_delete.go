package sensor

func (s *service) Delete(uuid string) (uint, error) {
	result, err := s.mgr.DeleteSensor(uuid)
	return result, err
}
