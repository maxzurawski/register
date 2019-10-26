package dbprovider

import (
	"testing"

	"github.com/maxzurawski/utilities/db"

	"github.com/maxzurawski/register/dto"

	"github.com/maxzurawski/register/model"
	"github.com/maxzurawski/utilities/symbols"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// suite struct
type UpdateAttributeSuite struct {
	suite.Suite
}

// init suite
func TestUpdateAttributeSuite(t *testing.T) {
	suite.Run(t, new(UpdateAttributeSuite))
}

// setup test function
func (ua *UpdateAttributeSuite) SetupTest() {
	EnvironmentPreparations()
}

// failure test in case when symbol is not set
func (ua *UpdateAttributeSuite) TestFailure_SymbolNotSet() {

	// Arrange.

	attributes, _ := Mgr.GetAttributes()
	attribute := (model.AttributeArray(attributes)).FilterBy(model.FilterBySymbol, symbols.AcceptableMin.String())
	assert.Equal(ua.T(), "ACCEPTABLE_MIN", *attribute.Symbol)
	dto := Mgr.MapAttributeToDTO(attribute)
	dto.Symbol = ""

	// Act.

	result, err := Mgr.UpdateAttribute(dto)

	// Assert.

	assert.NotNil(ua.T(), err)
	assert.Equal(ua.T(), "symbol is empty. nothing to update", err.Error())
	assert.Nil(ua.T(), result)
}

// failure test in case trying to update not existing attribute (with unknown symbol, like "REPORT_FREQUENCY")
func (ua *UpdateAttributeSuite) TestFailure_UpdateNotExistingAttribute() {

	// Arrange.

	dto := dto.AttributeDTO{
		Symbol:      "REPORT_FREQUENCY",
		Name:        "Frequency",
		Inputtype:   "numeric",
		Description: "Frequency attribute",
	}

	// Act.

	result, err := Mgr.UpdateAttribute(dto)

	// Assert.

	assert.NotNil(ua.T(), err)
	assert.Nil(ua.T(), result)
	assert.Equal(ua.T(), "record not found", err.Error())
}

// failure test in case trying to update existing attribute with changed inputtype
func (ua *UpdateAttributeSuite) TestFailure_InputtypeReassign() {

	// Arrange.

	attributes, _ := Mgr.GetAttributes()
	attribute := (model.AttributeArray(attributes)).FilterBy(model.FilterBySymbol, symbols.Active.String())
	assert.Equal(ua.T(), "boolean", *attribute.Inputtype)
	dto := Mgr.MapAttributeToDTO(attribute)
	dto.Inputtype = "numeric"

	// Act.

	result, err := Mgr.UpdateAttribute(dto)

	// Assert.

	assert.Nil(ua.T(), err)
	assert.NotNil(ua.T(), result.Inputtype)
	assert.Equal(ua.T(), "boolean", *result.Inputtype)
}

// failure test in case trying to update existing attribute with changed symbol
func (ua *UpdateAttributeSuite) TestFailure_SymbolReassign() {

	// Arrange.

	attributes, _ := Mgr.GetAttributes()
	attribute := (model.AttributeArray(attributes)).FilterBy(model.FilterBySymbol, symbols.AcceptableMax.String())
	assert.Equal(ua.T(), "ACCEPTABLE_MAX", *attribute.Symbol)
	dto := Mgr.MapAttributeToDTO(attribute)
	dto.Symbol = "FREQUENCY_OF_REPORTS" // NOTE - reassign

	// Act.

	result, err := Mgr.UpdateAttribute(dto)

	// Assert.

	assert.NotNil(ua.T(), err)
	assert.Equal(ua.T(), "record not found", err.Error())
	assert.Nil(ua.T(), result)
}

// success update with changing name and description
func (ua *UpdateAttributeSuite) TestSuccessUpdate() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	attributeToSave := model.Attribute{}
	symbol := "FREQUENCY_OF_REPORT"
	name := "Frequency"
	description := "Frequency description"
	inputtype := "numeric"

	attributeToSave.Symbol = &symbol
	attributeToSave.Name = &name
	attributeToSave.Description = &description
	attributeToSave.Inputtype = &inputtype

	Mgr.GetDb().Save(attributeToSave)

	attributes, _ := Mgr.GetAttributes()
	assert.Equal(ua.T(), 4, len(attributes))
	attribute := (model.AttributeArray(attributes)).FilterBy(model.FilterBySymbol, "FREQUENCY_OF_REPORT")
	assert.Equal(ua.T(), "FREQUENCY_OF_REPORT", *attribute.Symbol)
	assert.Equal(ua.T(), "Frequency", *attribute.Name)
	assert.Equal(ua.T(), "Frequency description", *attribute.Description)

	dto := Mgr.MapAttributeToDTO(attribute)

	dto.Name = "Frequency updated"
	dto.Description = "Frequency description updated"

	// Act.

	_, err := Mgr.UpdateAttribute(dto)

	// Assert.

	assert.Nil(ua.T(), err)

	afterUpdate, _ := Mgr.GetAttributes()
	assert.Equal(ua.T(), 4, len(attributes))
	attributeAfterUpdate := (model.AttributeArray(afterUpdate)).FilterBy(model.FilterBySymbol, "FREQUENCY_OF_REPORT")
	assert.Equal(ua.T(), "Frequency updated", *attributeAfterUpdate.Name)
	assert.Equal(ua.T(), "Frequency description updated", *attributeAfterUpdate.Description)
}
