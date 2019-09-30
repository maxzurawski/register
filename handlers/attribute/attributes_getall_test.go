package attribute

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xdevices/register/dto"

	"github.com/xdevices/register/model"

	"github.com/xdevices/register/services/attribute"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/xdevices/register/dbprovider"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleGetAllSuite struct {
	suite.Suite
}

// init suite
func TestHandleGetAllSuite(t *testing.T) {
	suite.Run(t, new(HandleGetAllSuite))
}

// setup test
func (h *HandleGetAllSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()

	// NOTE: we have to init our service - handlers use it
	attribute.Init()
}

// get when nothing found
func (h *HandleGetAllSuite) Test_nothing_found() {

	// Arrange.

	// NOTE: we have to delete all Attributes first - when initiating db manager we explicitly add attributes
	dbprovider.Mgr.GetDb().Delete(model.Attribute{})
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/attributes/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetAllAttributes(c)

	// Assert.

	assert.Equal(h.T(), http.StatusNoContent, rec.Code)
}

// get when data present
func (h *HandleGetAllSuite) Test_find_all() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/attributes/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetAllAttributes(c)

	// Assert.

	assert.Equal(h.T(), http.StatusOK, rec.Code)

	var attributes []dto.AttributeDTO
	_ = json.NewDecoder(rec.Body).Decode(&attributes)
	assert.Equal(h.T(), 3, len(attributes))
}
