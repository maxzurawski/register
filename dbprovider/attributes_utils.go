package dbprovider

import (
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/model"
)

func (mgr *manager) MapAttributeToDTO(attribute *model.Attribute) dto.AttributeDTO {

	result := dto.AttributeDTO{}

	if attribute == nil || attribute.Symbol == nil {
		return result
	}

	result.Name = *attribute.Name
	result.Symbol = *attribute.Symbol
	result.Description = *attribute.Description
	result.Inputtype = *attribute.Inputtype

	return result
}
