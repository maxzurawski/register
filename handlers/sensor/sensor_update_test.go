package sensor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/xdevices/register/dto"

	"github.com/xdevices/utilities/db"
	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/services/sensor"
)

// suite struct
type HandleSensorUpdateSuite struct {
	suite.Suite
}

// init suite
func TestHandleSensorUpdateSuite(t *testing.T) {
	suite.Run(t, new(HandleSensorUpdateSuite))
}

// setup test func
func (h *HandleSensorUpdateSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	sensor.Init()
}

// failure test - uuid empty
func (h *HandleSensorUpdateSuite) TestFailure_empty_uuid() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "uuid parameter cannot be empty", err.Msg)
}

// failure test - uuid not valid
func (h *HandleSensorUpdateSuite) TestFailure_uuid_invalid() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("not-valid-uuid")

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "invalid uuid given", err.Msg)
}

// failure test - unsupported media type
func (h *HandleSensorUpdateSuite) TestFailure_body_unsupported_media_type() {

	// Arrange.

	var sensorRegisters []dto.SensorRegisterDTO
	sensorRegisters = append(sensorRegisters, dto.SensorRegisterDTO{})
	bytes, _ := json.Marshal(sensorRegisters)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", strings.NewReader(string(bytes)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusUnsupportedMediaType, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "code=415, message=Unsupported Media Type", err.Msg)
}

// failure test - body not of type dto.SensorRegisterDTO
func (h *HandleSensorUpdateSuite) TestFailure_body_invalid() {

	// Arrange.

	var sensorRegisters []dto.SensorRegisterDTO
	sensorRegisters = append(sensorRegisters, dto.SensorRegisterDTO{})
	bytes, _ := json.Marshal(sensorRegisters)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "code=400, message=Unmarshal type error: expected=dto.SensorRegisterDTO, got=array, field=, offset=1", err.Msg)
}

// failure test - versions are different
func (h *HandleSensorUpdateSuite) TestFailure_versions_difference() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	prepareTestSensor()
	registerDTO2Update := sensor.Service.FindSensorByUuid("81750491-88dd-410e-b53f-1666786cd721")
	registerDTO2Update.Version = 1
	bytes, _ := json.Marshal(registerDTO2Update)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(registerDTO2Update.Uuid)

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "optimistic lock exception. versions are different", err.Msg)
}

// success test
func (h *HandleSensorUpdateSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	prepareTestSensor()
	registerDTO2Update := sensor.Service.FindSensorByUuid("81750491-88dd-410e-b53f-1666786cd721")
	description4Rollback := registerDTO2Update.Description
	registerDTO2Update.Description = registerDTO2Update.Description + " Updated"
	bytes, _ := json.Marshal(registerDTO2Update)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:uuid", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues(registerDTO2Update.Uuid)

	// Act.

	_ = HandleSensorUpdate(c)

	// Assert.

	assert.Equal(h.T(), http.StatusAccepted, rec.Code)
	var afterUpdate dto.SensorRegisterDTO
	_ = json.Unmarshal(rec.Body.Bytes(), &afterUpdate)
	assert.True(h.T(), registerDTO2Update.Version < afterUpdate.Version)
	assert.True(h.T(), registerDTO2Update.ID == afterUpdate.ID)
	assert.Equal(h.T(), description4Rollback+" Updated", afterUpdate.Description)
}
