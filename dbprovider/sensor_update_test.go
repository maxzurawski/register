package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maxzurawski/register/model"

	"github.com/maxzurawski/utilities/db"
	"github.com/maxzurawski/utilities/symbols"
	"github.com/stretchr/testify/suite"
)

// suite struct
type SensorUpdateSuite struct {
	suite.Suite
}

// init suite
func TestSensorUpdateSuite(t *testing.T) {
	suite.Run(t, new(SensorUpdateSuite))
}

// setup test
func (s *SensorUpdateSuite) SetupTest() {
	EnvironmentPreparations()
}

// check updating version and created at field in attribute
func (s *SensorUpdateSuite) TestUpdatingVersionAndCreatedAtField() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareFakeSensorRegisters()
	register, _ := Mgr.GetSensorByUuid("b255382c-35c8-4aee-99db-74fbe41ca9f7")
	attribute := (model.SensorAttributes(register.Attributes)).FilterBy(filterSensorAttributeBySymbol, symbols.Active.String())

	createdAt := *attribute.CreateAt
	version := *attribute.Version
	value := *attribute.Value
	id := *attribute.ID

	*attribute.Value = "false"

	dto := Mgr.MapToSensorDTO(register)

	// Act.

	result, _ := Mgr.UpdateSensor(dto)

	// Assert.

	assert.NotNil(s.T(), result)

	attributeAfterUpdate := (model.SensorAttributes(result.Attributes)).FilterBy(filterSensorAttributeBySymbol, symbols.Active.String())
	assert.Equal(s.T(), id, *attributeAfterUpdate.ID)
	assert.True(s.T(), version+1 == *attributeAfterUpdate.Version)
	assert.True(s.T(), value != *attributeAfterUpdate.Value)
	assert.True(s.T(), createdAt.Unix() == attributeAfterUpdate.CreateAt.Unix())
}
