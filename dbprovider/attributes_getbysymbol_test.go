package dbprovider

import (
	"testing"

	"github.com/maxzurawski/utilities/symbols"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

// suite struct
type GetBySymbolSuite struct {
	suite.Suite
}

// init suite
func TestGetBySymbolSuite(t *testing.T) {
	suite.Run(t, new(GetBySymbolSuite))
}

// setup test function
func (g *GetBySymbolSuite) SetupTest() {
	EnvironmentPreparations()
}

// test failure - by providing falsy symbol, like "FREQUENCY_OF_REPORT
func (g *GetBySymbolSuite) TestFailure_unknown_symbol() {

	// Arrange && Act.

	result, err := Mgr.GetAttributeBySymbol("FREQUENCY_OF_REPORT")

	// Assert.

	assert.Nil(g.T(), result)
	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "record not found", err.Error())
}

// test success - by providing AcceptableMax.String value
func (g *GetBySymbolSuite) TestSuccess() {

	// Arrange & Act.

	result, err := Mgr.GetAttributeBySymbol(symbols.AcceptableMin.String())

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), symbols.AcceptableMin.String(), *result.Symbol)
	assert.Equal(g.T(), "Acceptable minimum value", *result.Name)
	assert.Equal(g.T(), "Minimum value acceptable, before notification happens", *result.Description)
	assert.Equal(g.T(), "numeric", *result.Inputtype)
}
