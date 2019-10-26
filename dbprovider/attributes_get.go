package dbprovider

import "github.com/maxzurawski/register/model"

func (mgr *manager) GetAttributes() ([]model.Attribute, error) {
	var attributes []model.Attribute
	err := mgr.GetDb().Find(&attributes).Error
	return attributes, err
}
