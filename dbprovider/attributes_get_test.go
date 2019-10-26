package dbprovider

import (
	"testing"

	"github.com/xdevices/utilities/symbols"

	"github.com/xdevices/register/model"

	"github.com/stretchr/testify/assert"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// GetAttributesTestSuite struct
type GetAttributesTestSuite struct {
	suite.Suite
}

// SetupTest
func (g *GetAttributesTestSuite) SetupTest() {
	EnvironmentPreparations()
}

// TestGetAttributesTestSuite
func TestGetAttributesTestSuite(t *testing.T) {
	suite.Run(t, new(GetAttributesTestSuite))
}

// Test if attributes are present
func (g *GetAttributesTestSuite) TestGetAttributes_Succes() {

	// Arrange.

	// REMEMBER to use Cleanup hook
	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	// Act.

	result, err := Mgr.GetAttributes()

	// Assert.

	assert.Nil(g.T(), err)
	assert.Equal(g.T(), 3, len(result))

	attribute := (model.AttributeArray(result)).FilterBy(model.FilterBySymbol, symbols.AcceptableMax.String())
	assert.Equal(g.T(), "Acceptable maximum value", *attribute.Name)
	assert.Equal(g.T(), "Maximum value acceptable, before notification happens", *attribute.Description)
	assert.Equal(g.T(), "numeric", *attribute.Inputtype)

	attribute = (model.AttributeArray(result)).FilterBy(model.FilterBySymbol, symbols.Active.String())
	assert.Equal(g.T(), "Active flag", *attribute.Name)
	assert.Equal(g.T(), "Is sensor active, or should it be ignored?", *attribute.Description)
	assert.Equal(g.T(), "boolean", *attribute.Inputtype)
}
