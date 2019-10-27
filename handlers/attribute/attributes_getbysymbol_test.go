package attribute

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxzurawski/register/dto"

	"github.com/labstack/echo"

	"github.com/maxzurawski/register/dbprovider"
	"github.com/maxzurawski/register/services/attribute"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleGetBySymbolSuite struct {
	suite.Suite
}

// init suite
func TestHandleGetBySymbolSuite(t *testing.T) {
	suite.Run(t, new(HandleGetBySymbolSuite))
}

// setup test func
func (h *HandleGetBySymbolSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	attribute.Init()
}

// get by symbol unknown
func (h *HandleGetBySymbolSuite) Test_get_unknown() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/attributes/:symbol", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("FREQUENCY_OF_REPORT")

	// Act.

	_ = HandleGetAttributeBySymbol(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
}

// get by symbol known
func (h *HandleGetBySymbolSuite) Test_get_known() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/attributes/:symbol", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:symbol")
	c.SetParamNames("symbol")
	c.SetParamValues("ACTIVE")

	// Act.

	_ = HandleGetAttributeBySymbol(c)

	// Assert.

	assert.Equal(h.T(), http.StatusOK, rec.Code)
	dto := dto.AttributeDTO{}
	_ = json.NewDecoder(rec.Body).Decode(&dto)

	assert.Equal(h.T(), "Active flag", dto.Name)
	assert.Equal(h.T(), "Is sensor active, or should it be ignored?", dto.Description)
	assert.Equal(h.T(), "boolean", dto.Inputtype)
	assert.Equal(h.T(), "ACTIVE", dto.Symbol)
}
