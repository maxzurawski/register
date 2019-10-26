package dbprovider

import (
	"errors"

	"github.com/maxzurawski/register/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) GetSensorByUuid(uuid string) (*model.SensorRegister, error) {

	// check if uuid given is not empty
	if stringutils.IsZero(uuid) {
		return nil, errors.New("given uuid is empty. nothing to search for")
	}

	// check if uuid is valid
	if !stringutils.IsUuidValid(uuid) {
		return nil, errors.New("given uuid is not valid")
	}

	// query for sensor
	sensorRegister := &model.SensorRegister{}
	err := mgr.GetDb().Where("uuid=?", uuid).Preload("Attributes.Attribute").First(sensorRegister).Error
	if err != nil {
		return nil, err
	}
	return sensorRegister, nil
}
