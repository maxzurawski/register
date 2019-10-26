package sensor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"

	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/services/sensor"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleSensorDeleteSuite struct {
	suite.Suite
}

// init suite
func TestHandleSensorDeleteSuite(t *testing.T) {
	suite.Run(t, new(HandleSensorDeleteSuite))
}

// setup test func
func (h *HandleSensorDeleteSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	sensor.Init()
}

// failure test - uuid empty
func (h *HandleSensorDeleteSuite) TestFailure_uuid_empty() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:uuid", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleSensorDelete(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "uuid parameter cannot be empty", err.Msg)
}

// failure test - uuid invalid
func (h *HandleSensorDeleteSuite) TestFailure_uuid_invalid() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:uuid", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("not-valid-uuid")

	// Act.

	_ = HandleSensorDelete(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "invalid uuid given", err.Msg)
}

// failure test - sensor does not exists
func (h *HandleSensorDeleteSuite) TestFailure_sensor_not_exists() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:uuid", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleSensorDelete(c)

	// Assert.

	assert.Equal(h.T(), http.StatusNotFound, rec.Code)
}

// success test
func (h *HandleSensorDeleteSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()

	prepareTestSensor()

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:uuid", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleSensorDelete(c)

	// Assert.

	assert.Equal(h.T(), http.StatusOK, rec.Code)
}
