package dbprovider

import (
	"errors"

	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) UpdateAttribute(attributeDTO dto.AttributeDTO) (*model.Attribute, error) {

	// check if attributeDTO.Symbol is not zero
	if stringutils.IsZero(attributeDTO.Symbol) {
		return nil, errors.New("symbol is empty. nothing to update")
	}

	// check if attributeDTO with Symbol exists at all in our database
	attribute := &model.Attribute{}
	err := mgr.GetDb().Where("symbol=?", attributeDTO.Symbol).First(attribute).Error
	if err != nil {
		return nil, err
	}

	// reassing name and descritpion ONLY
	*attribute.Name = attributeDTO.Name
	*attribute.Description = attributeDTO.Description
	err = mgr.GetDb().Save(attribute).Error
	if err != nil {
		return nil, err
	}
	return attribute, nil
}
