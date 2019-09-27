package dbprovider

import (
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/model"
	"github.com/xdevices/utilities/stringutils"
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

	err := mgr.GetDb().Unscoped().Where("sensor_register_id=?", *register.ID).Delete(model.SensorAttribute{}).Error
	if err != nil {
		log.Warn(err)
	}

	// reset sensorattributes on object -> delete current sensorattributes, assign sensorattributes from dto
	register.Attributes = []model.SensorAttribute{}
	register.Attributes = convert(registerDTO.Attributes, modificationTime, register)
}

func convert(attributesDTO []dto.SensorAttributeDTO, now time.Time, register *model.SensorRegister) (attributes []model.SensorAttribute) {

	attrById := map[uint]model.SensorAttribute{}
	for _, attr := range register.Attributes {
		attrById[*attr.ID] = attr
	}

	for _, item := range attributesDTO {

		sensorAttribute := new(model.SensorAttribute)
		// NOTE: This is very tricky. Had to init all pointer fields, otherwise it was using reference
		// in each added item in attributes.
		sensorAttribute.RefSymbol = new(string)
		sensorAttribute.Value = new(string)
		sensorAttribute.Version = new(uint)
		sensorAttribute.CreateAt = new(time.Time)
		sensorAttribute.ModifiedAt = new(time.Time)

		if !stringutils.IsZero(item.ID) {
			sensorAttribute.ID = new(uint)
			*sensorAttribute.ID = item.ID
		}

		*sensorAttribute.RefSymbol = item.Symbol
		*sensorAttribute.Value = item.Value.ToString()

		if attribute, ok := attrById[item.ID]; !ok {
			*sensorAttribute.CreateAt = now
			*sensorAttribute.Version = item.Version
		} else {
			*sensorAttribute.CreateAt = *attribute.CreateAt
			*sensorAttribute.Version = *attribute.Version + 1
		}
		*sensorAttribute.ModifiedAt = now
		attributes = append(attributes, *sensorAttribute)
	}
	return
}
