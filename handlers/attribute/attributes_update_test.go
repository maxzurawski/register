package attribute

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/xdevices/utilities/resterror"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
	"github.com/xdevices/register/dto"

	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/services/attribute"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleAttributeUpdate struct {
	suite.Suite
}

// init suite
func TestHandleAttributeUpdate(t *testing.T) {
	suite.Run(t, new(HandleAttributeUpdate))
}

// setup test func
func (h *HandleAttributeUpdate) SetupTest() {
	dbprovider.EnvironmentPreparations()
	attribute.Init()
}

// test symbol empty
func (h *HandleAttributeUpdate) Test_symbol_empty() {

	// Arrange.

	inputDTO := dto.AttributeDTO{
		Symbol:      "ACTIVE",
		Name:        "Active flag",
		Description: "Is sensor active, or should it be ignored?",
		Inputtype:   "boolean",
	}
	bytes, _ := json.Marshal(inputDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:symbol", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("")

	// Act.

	_ = HandleUpdateAttribute(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.NewDecoder(rec.Body).Decode(&err)
	assert.Equal(h.T(), "not known symbol", err.Msg)
}

// test post body incorrect
func (h *HandleAttributeUpdate) Test_body_incorrect() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:symbol", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("ACTIVE")

	// Act.

	_ = HandleUpdateAttribute(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.NewDecoder(rec.Body).Decode(&err)
	assert.Equal(h.T(), "code=400, message=Request body can't be empty", err.Msg)
}

// test symbol in body empty
func (h *HandleAttributeUpdate) Test_symbol_in_body_empty() {

	// Arrange.

	inputDTO := dto.AttributeDTO{
		Symbol:      "",
		Name:        "Active flag",
		Description: "Is sensor active, or should it be ignored?",
		Inputtype:   "boolean",
	}
	bytes, _ := json.Marshal(inputDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:symbol", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("ACTIVE")

	// Act.

	_ = HandleUpdateAttribute(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.NewDecoder(rec.Body).Decode(&err)
	assert.Equal(h.T(), "given symbol of attribute cannot be zero or empty", err.Msg)
}

// test symbol given in path does not match symbol in body
func (h *HandleAttributeUpdate) Test_symbol_in_body_not_match() {

	// Arrange.

	inputDTO := dto.AttributeDTO{
		Symbol: "REPORT_OF_FREQUENCY",
	}
	bytes, _ := json.Marshal(inputDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:symbol", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("ACTIVE")

	// Act.

	_ = HandleUpdateAttribute(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.NewDecoder(rec.Body).Decode(&err)
	assert.Equal(h.T(), "given symbol with symbol in attribute provided does not match", err.Msg)
}

// test status accepted - change description - and find changed attribute, revert changes by resetting description to original
func (h *HandleAttributeUpdate) Test_Update_Accepted() {

	// Arrange.

	attributeDTO, _ := attribute.Service.GetAttributeBySymbol("ACTIVE")
	assert.Equal(h.T(), "Is sensor active, or should it be ignored?", attributeDTO.Description)
	attributeDTO.Description = "Is sensor active, or should it be ignored? Updated"
	bytes, _ := json.Marshal(attributeDTO)

	// "Is sensor active, or should it be ignored?"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:symbol", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("ACTIVE")

	// Act.

	_ = HandleUpdateAttribute(c)

	// Assert.

	assert.Equal(h.T(), http.StatusAccepted, rec.Code)
	var attributeAfterUpdate dto.AttributeDTO
	_ = json.NewDecoder(rec.Body).Decode(&attributeAfterUpdate)
	assert.Equal(h.T(), "Is sensor active, or should it be ignored? Updated", attributeAfterUpdate.Description)

	// Reset description to original
	attributeAfterUpdate.Description = "Is sensor active, or should it be ignored?"
	_, _ = attribute.Service.UpdateAttribute(attributeAfterUpdate)

}
