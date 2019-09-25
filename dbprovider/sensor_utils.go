package dbprovider

import (
	"time"

	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/model"
)

func (mgr *manager) MapToSensorEntity(registerDTO dto.SensorRegisterDTO) *model.SensorRegister {

	now := time.Now()
	version := uint(0)

	sensor := &model.SensorRegister{
		Name:        &registerDTO.Name,
		Version:     &version,
		Uuid:        &registerDTO.Uuid,
		Type:        &registerDTO.Type,
		Description: &registerDTO.Description,
		Attributes:  mgr.MapToSensorAttribute(registerDTO.Attributes, now),
		CreatedAt:   &now,
		ModifiedAt:  &now,
	}

	return sensor
}

func (mgr *manager) MapToSensorDTO(sensor *model.SensorRegister) dto.SensorRegisterDTO {

	if sensor == nil {
		return dto.SensorRegisterDTO{}
	}

	var attr []dto.SensorAttributeDTO

	for _, item := range sensor.Attributes {
		attr = append(attr, mgr.MapToSensorAttributeDTO(&item))
	}

	dto := dto.SensorRegisterDTO{
		ID:          *sensor.ID,
		Version:     *sensor.Version,
		Name:        *sensor.Name,
		Type:        *sensor.Type,
		Description: *sensor.Description,
		Uuid:        *sensor.Uuid,
		Attributes:  attr,
	}

	return dto
}
