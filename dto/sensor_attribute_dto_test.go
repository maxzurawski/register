package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSensorAttributeMarshalling(t *testing.T) {

	// Arrange.

	dto := SensorAttributeDTO{
		Symbol:  "TEST",
		Value:   "false",
		ID:      1,
		Version: 0,
	}

	// Act.

	bytes, err := json.Marshal(dto)

	// Assert.

	assert.Nil(t, err)
	assert.Equal(t, `{"id":1,"version":0,"symbol":"TEST","value":false,"sensor_id":0}`, string(bytes))
}

func TestReadJson(t *testing.T) {

	// Arrange.

	input := `{"id":1,"version":0,"symbol":"TEST","value":true,"sensor_id":0}`
	dto := SensorAttributeDTO{}

	// Act.

	err := json.Unmarshal([]byte(input), &dto)

	// Assert.

	assert.Nil(t, err)
	assert.Equal(t, "true", dto.Value.ToString())
}
