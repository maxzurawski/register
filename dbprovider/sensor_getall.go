package dbprovider

import "github.com/maxzurawski/register/model"

func (mgr *manager) GetAllSensors() ([]model.SensorRegister, error) {
	var sensors []model.SensorRegister
	err := mgr.GetDb().Preload("Attributes.Attribute").Find(&sensors).Error
	if err != nil {
		return nil, err
	}
	return sensors, nil
}
