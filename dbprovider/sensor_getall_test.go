package dbprovider

import (
	"strings"
	"testing"

	"github.com/xdevices/utilities/symbols"

	"github.com/xdevices/register/model"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// suite struct
type GetAllSuite struct {
	suite.Suite
}

// init suite
func TestGetAllSuite(t *testing.T) {
	suite.Run(t, new(GetAllSuite))
}

// setup test
func (g *GetAllSuite) SetupTest() {
	EnvironmentPreparations()
}

// failure test - when nothing is there // NOTE: kind'a success search, because nothing is stored
func (g *GetAllSuite) TestFailure_NothingFound() {

	// Arrange & Act.

	result, err := Mgr.GetAllSensors()

	// Assert.
	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), 0, len(result))
}

// success test - prepare sample sensor // REMEMBER: use cleaner hook
func (g *GetAllSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareFakeSensorRegisters()

	// Act.

	result, err := Mgr.GetAllSensors()

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), 2, len(result))

	sensor := (model.SensorsArray(result)).FilterBy(filterSensorByUuid, "a607f675-e106-4617-9261-bcd5ee5f2ad7")
	assert.NotNil(g.T(), sensor)
	assert.True(g.T(), *sensor.ID > 0)
	assert.Equal(g.T(), uint(2), *sensor.ID)
	assert.Equal(g.T(), "test 2", *sensor.Name)
	assert.Equal(g.T(), "test 2 description", *sensor.Description)
	assert.Equal(g.T(), "TEST2_TYPE", *sensor.Type)
	assert.Equal(g.T(), 2, len(sensor.Attributes))

	attributes := model.SensorAttributes(sensor.Attributes)
	attribute := attributes.FilterBy(filterSensorAttributeBySymbol, symbols.Active.String())
	assert.NotNil(g.T(), attribute)
	assert.True(g.T(), *attribute.ID > 0)
	assert.True(g.T(), *attribute.SensorRegisterID == *sensor.ID)
	assert.Equal(g.T(), "true", *attribute.Value)

	attribute = attributes.FilterBy(filterSensorAttributeBySymbol, symbols.AcceptableMin.String())
	assert.Equal(g.T(), "24.5", *attribute.Value)

}

func filterSensorByUuid(sensor *model.SensorRegister, uid string) bool {
	if *sensor.Uuid == uid {
		return true
	}
	return false
}

func filterSensorAttributeBySymbol(attribute *model.SensorAttribute, symbol string) bool {
	if *attribute.RefSymbol == strings.ToUpper(symbol) {
		return true
	}
	return false
}
