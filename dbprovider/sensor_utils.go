package dbprovider

import (
	"time"

	"github.com/maxzurawski/utilities/symbols"

	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/model"
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

func (mgr *manager) MapToCachedSensorDTO(sensor *model.SensorRegister) dto.CachedSensorDTO {

	if sensor == nil {
		return dto.CachedSensorDTO{}
	}

	dto := dto.CachedSensorDTO{
		Name:        *sensor.Name,
		Type:        *sensor.Type,
		Description: *sensor.Description,
		Uuid:        *sensor.Uuid,
		Active:      sensor.IsActive(),
		Max:         sensor.GetAttributeAsString(symbols.AcceptableMax.String()),
		Min:         sensor.GetAttributeAsString(symbols.AcceptableMin.String()),
		Nacta:       sensor.GetAttributeAsInt(symbols.NotificationAfterContinuousTransitionAmount.String()),
	}

	return dto
}
