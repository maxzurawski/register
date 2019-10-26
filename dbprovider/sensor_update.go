package dbprovider

import (
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) UpdateSensor(sensorDTO dto.SensorRegisterDTO) (*model.SensorRegister, error) {
	// check if ID is not 0
	if stringutils.IsZero(sensorDTO.ID) {
		return nil, errors.New(fmt.Sprintf("given sensor does not have valid id. nothing to update"))
	}

	// check if uuid is valid
	if !stringutils.IsUuidValid(sensorDTO.Uuid) {
		return nil, errors.New("invalid uuid given")
	}

	// check if when searching for the sensor with id given in dto there will be no error
	register := &model.SensorRegister{}
	err := mgr.GetDb().Where("ID=?", sensorDTO.ID).Preload("Attributes.Attribute").First(register).Error
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	// check if register with ID exists
	if register.ID == nil {
		return nil, errors.New(fmt.Sprintf("sensor with id [%d] does not exists. nothing to update", sensorDTO.ID))
	}

	// check if found register has the same version as given dto.version
	version := *register.Version
	if version != sensorDTO.Version {
		return nil, errors.New("entitie's version differs from version of given input. please update your state")
	}

	// pull up version
	version = version + 1

	// set modification time to now
	now := time.Now()
	register.Version = &version
	register.ModifiedAt = &now

	// populate changes on sensor register object
	mgr.populateWithChanges(register, sensorDTO, now)

	// save changes
	err = mgr.db.Save(register).Error
	return register, err
}

func (mgr *manager) populateWithChanges(register *model.SensorRegister, registerDTO dto.SensorRegisterDTO, modificationTime time.Time) {

	register.Uuid = &registerDTO.Uuid
	register.Type = &registerDTO.Type
	register.Name = &registerDTO.Name
	register.Description = &registerDTO.Description

	createdTimeById := make(map[uint]time.Time)
	for _, attr := range register.Attributes {
		createdTimeById[*attr.ID] = *attr.CreateAt
	}

	err := mgr.GetDb().Unscoped().Where("sensor_register_id=?", *register.ID).Delete(model.SensorAttribute{}).Error
	if err != nil {
		log.Warn(err)
	}

	// reset sensorattributes on object -> delete current sensorattributes, assign sensorattributes from dto
	register.Attributes = []model.SensorAttribute{}
	register.Attributes = convert(registerDTO.Attributes, modificationTime, createdTimeById)
}

func convert(attributesDTO []dto.SensorAttributeDTO, now time.Time, createdTimeById map[uint]time.Time) (attributes []model.SensorAttribute) {

	for _, item := range attributesDTO {

		sensorAttribute := new(model.SensorAttribute)
		// NOTE: This is very tricky. Had to init all pointer fields, otherwise it was using reference
		// in each added item in attributes.
		sensorAttribute.RefSymbol = new(string)
		sensorAttribute.Value = new(string)
		sensorAttribute.Version = new(uint)
		sensorAttribute.CreateAt = new(time.Time)
		sensorAttribute.ModifiedAt = new(time.Time)

		*sensorAttribute.RefSymbol = item.Symbol
		*sensorAttribute.Value = item.Value.ToString()
		*sensorAttribute.CreateAt = now
		*sensorAttribute.ModifiedAt = now

		if !stringutils.IsZero(item.ID) {
			sensorAttribute.ID = new(uint)
			*sensorAttribute.ID = item.ID
			*sensorAttribute.Version = item.Version + 1
			*sensorAttribute.CreateAt = createdTimeById[item.ID]
		}
		attributes = append(attributes, *sensorAttribute)
	}
	return
}
