package attribute

import "github.com/xdevices/register/dto"

func (s *service) UpdateAttribute(attribute dto.AttributeDTO) (*dto.AttributeDTO, error) {
	result, err := s.mgr.UpdateAttribute(attribute)
	dto := s.mgr.MapAttributeToDTO(result)
	return &dto, err
}
