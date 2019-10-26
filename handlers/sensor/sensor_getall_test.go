package sensor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maxzurawski/utilities/db"

	"github.com/labstack/echo"

	"github.com/maxzurawski/register/dbprovider"
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/services/sensor"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleSensorsGetAllSuite struct {
	suite.Suite
}

// init suite
func TestHandleSensorsGetAllSuite(t *testing.T) {
	suite.Run(t, new(HandleSensorsGetAllSuite))
}

// setup test func
func (h *HandleSensorsGetAllSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	sensor.Init()
}

// failure test - no content
func (h *HandleSensorsGetAllSuite) TestFailure_getall_no_content() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetSensors(c)

	// Assert.

	assert.Equal(h.T(), http.StatusNoContent, rec.Code)
}

// success test - content available
func (h *HandleSensorsGetAllSuite) TestSuccess_getall_content_available() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()

	prepareTestSensor()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetSensors(c)

	// Assert.
	assert.Equal(h.T(), http.StatusOK, rec.Code)
	var dtos []dto.SensorRegisterDTO
	_ = json.NewDecoder(rec.Body).Decode(&dtos)
	assert.Equal(h.T(), 1, len(dtos))
	afterGetAll := dtos[0]
	assert.Equal(h.T(), "Dummy sensor", afterGetAll.Name)
	assert.Equal(h.T(), "81750491-88dd-410e-b53f-1666786cd721", afterGetAll.Uuid)
	assert.True(h.T(), afterGetAll.ID > 0)
}
