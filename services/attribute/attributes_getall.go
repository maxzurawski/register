package attribute

import "github.com/xdevices/register/dto"

func (s *service) GetAll() ([]dto.AttributeDTO, error) {
	attributes, err := s.mgr.GetAttributes()
	if err != nil {
		return nil, err
	}

	var dtos []dto.AttributeDTO
	if len(attributes) == 0 {
		return dtos, nil
	}

	for _, item := range attributes {
		dtos = append(dtos, s.mgr.MapAttributeToDTO(&item))
	}
	return dtos, nil
}
