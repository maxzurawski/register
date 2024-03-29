package sensor

import (
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/services/sensor"
	"github.com/maxzurawski/utilities/symbols"
)

func prepareTestSensor() {
	var attributes []dto.SensorAttributeDTO
	attributes = append(attributes, dto.SensorAttributeDTO{Symbol: symbols.Active.String(), Value: "true"})

	registerDTO := dto.SensorRegisterDTO{
		Name:       "Dummy sensor",
		Type:       "DUMMY_TYPE",
		Uuid:       "81750491-88dd-410e-b53f-1666786cd721",
		Attributes: attributes,
	}
	_, _ = sensor.Service.Save(registerDTO)
}
