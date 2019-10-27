package dbprovider

import (
	"strings"
	"testing"

	"github.com/maxzurawski/utilities/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// suite struct
type GetByUuidSuite struct {
	suite.Suite
}

// init suite
func TestGetByUuidSuite(t *testing.T) {
	suite.Run(t, new(GetByUuidSuite))
}

// setup test function
func (g *GetByUuidSuite) SetupTest() {
	EnvironmentPreparations()
}

// failure check - look for empty-string-uuid
func (g *GetByUuidSuite) TestFailure_EmptyStringUuid() {

	// Arrange & Act.

	result, err := Mgr.GetSensorByUuid("")

	// Assert.

	assert.Nil(g.T(), result)
	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "given uuid is empty. nothing to search for", err.Error())
}

// failure check - not valid uuid as parameter
func (g *GetByUuidSuite) TestFailure_NotValidUuid() {

	// Arrange & Act.

	result, err := Mgr.GetSensorByUuid("not-valid_uuid")

	// Assert.

	assert.Nil(g.T(), result)
	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "given uuid is not valid", err.Error())
}

// failure check - look for unknown uuid
func (g *GetByUuidSuite) TestFailure_UnkownUuid() {

	// Arrange & Act.

	result, err := Mgr.GetSensorByUuid("b255382c-35c8-4aee-99db-74fbe41ca9f7")

	// Assert.

	assert.Nil(g.T(), result)
	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "record not found", err.Error())
}

// success check - look for known uuid // REMEMBER to use cleaner hook
func (g *GetByUuidSuite) TestSucces() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareFakeSensorRegisters()

	// Act.

	result, err := Mgr.GetSensorByUuid("b255382c-35c8-4aee-99db-74fbe41ca9f7")

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), uint(1), *result.ID)
	assert.Equal(g.T(), "test 1", *result.Name)
	assert.Equal(g.T(), strings.ToUpper("dummy_type"), *result.Type)
}
