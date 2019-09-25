package dbprovider

import (
	"errors"
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/model"
	"github.com/xdevices/utilities/stringutils"
)

func (mgr *manager) SaveSensor(sensor dto.SensorRegisterDTO) (*model.SensorRegister, error) {

	// check if given sensor.ID is not zero -> if it set, fail
	if !stringutils.IsZero(sensor.ID) {
		err := fmt.Sprintf("sensor with given id exists. no save. try to update")
		log.Error(err)
		return nil, errors.New(err)
	}

	// check if uuid is valid
	valid := stringutils.IsUuidValid(sensor.Uuid)
	if !valid {
		return nil, errors.New("uuid is not valid")
	}

	// check if sensor with sensor.Uuid does not exists -> if there is an error other then "record not found"
	// it means, we problems with database
	entity := &model.SensorRegister{}

	err := mgr.GetDb().Where("uuid=?", sensor.Uuid).First(&entity).Error
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	// check if sensor with sensor.Uuid has an ID -> if it has, no save
	if entity.ID != nil {
		err := fmt.Sprintf("ups, something went wrong. sensor regisiter with uuid=[%s] already exists", sensor.Uuid)
		log.Error(err)
		return nil, errors.New(err)
	}

	// convert given dto to entity
	sensorRegister := mgr.MapToSensorEntity(sensor)

	// use create method to create register
	err = mgr.GetDb().Create(sensorRegister).Error

	if err != nil {
		return nil, err
	}

	// reload sensor register with all associated elements and return it back
	err = mgr.GetDb().Where("uuid=?", sensor.Uuid).Preload("Attributes.Attribute").Find(&sensorRegister).Error
	return sensorRegister, err

}
