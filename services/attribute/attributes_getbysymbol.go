package attribute

import "github.com/maxzurawski/register/dto"

func (s *service) GetAttributeBySymbol(symbol string) (*dto.AttributeDTO, error) {
	result, err := s.mgr.GetAttributeBySymbol(symbol)
	dto := s.mgr.MapAttributeToDTO(result)
	return &dto, err
}
