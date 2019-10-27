package dbprovider

import "github.com/maxzurawski/register/model"

func (mgr *manager) GetAttributeBySymbol(symbol string) (*model.Attribute, error) {
	attribute := &model.Attribute{}
	err := mgr.GetDb().Where("symbol=?", symbol).First(attribute).Error
	if err != nil {
		return nil, err
	}
	return attribute, nil
}
