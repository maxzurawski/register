package dbprovider

import (
	"time"

	"github.com/xdevices/utilities/stringutils"

	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/model"
)

func (mgr *manager) MapToSensorAttribute(attributesDTO []dto.SensorAttributeDTO, now time.Time) (attributes []model.SensorAttribute) {

	if attributesDTO == nil || len(attributesDTO) == 0 {
		return nil
	}

	for _, item := range attributesDTO {
		sensorAttribute := new(model.SensorAttribute)
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
		*sensorAttribute.Version = item.Version + 1
		*sensorAttribute.CreateAt = now
		*sensorAttribute.ModifiedAt = now
		attributes = append(attributes, *sensorAttribute)
	}
	return
}

func (mgr *manager) MapToSensorAttributeDTO(entity *model.SensorAttribute) dto.SensorAttributeDTO {

	dto := dto.SensorAttributeDTO{
		Version:          *entity.Version,
		ID:               *entity.ID,
		Value:            stringutils.ToMultiString(*entity.Value),
		Symbol:           *entity.RefSymbol,
		SensorRegisterID: *entity.SensorRegisterID,
	}
	return dto
}
