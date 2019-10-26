package sensor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xdevices/register/dto"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"

	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/services/sensor"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleGetSensorByUuidSuite struct {
	suite.Suite
}

// init suite
func TestHandleGetSensorByUuidSuite(t *testing.T) {
	suite.Run(t, new(HandleGetSensorByUuidSuite))
}

// setup test func
func (h *HandleGetSensorByUuidSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	sensor.Init()
}

// failure - empty uuid
func (h *HandleGetSensorByUuidSuite) TestFailure_uuid_empty() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetSensorByUuid(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "uuid parameter cannot be empty", err.Msg)
}

// failure - invalid uuid
func (h *HandleGetSensorByUuidSuite) TestFailure_uuidInvalid() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("not-valid-uuid")

	// Act.

	_ = HandleGetSensorByUuid(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "invalid uuid given", err.Msg)
}

// failure - no content
func (h *HandleGetSensorByUuidSuite) Test_no_content() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleGetSensorByUuid(c)

	// Assert.

	assert.Equal(h.T(), http.StatusNoContent, rec.Code)
	var err resterror.ErrorMsg
	_ = json.Unmarshal(rec.Body.Bytes(), &err)
	assert.Equal(h.T(), "unknown uuid given", err.Msg)
}

// success
func (h *HandleGetSensorByUuidSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	prepareTestSensor()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:uuid", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath("/:uuid")
	c.SetParamNames("uuid")
	c.SetParamValues("81750491-88dd-410e-b53f-1666786cd721")

	// Act.

	_ = HandleGetSensorByUuid(c)

	// Assert.

	assert.Equal(h.T(), http.StatusOK, rec.Code)
	var registerDTO dto.SensorRegisterDTO
	_ = json.Unmarshal(rec.Body.Bytes(), &registerDTO)
	assert.Equal(h.T(), "Dummy sensor", registerDTO.Name)
	assert.Equal(h.T(), "DUMMY_TYPE", registerDTO.Type)
	assert.Equal(h.T(), 1, len(registerDTO.Attributes))
	assert.True(h.T(), registerDTO.ID > 0)
	assert.Equal(h.T(), "81750491-88dd-410e-b53f-1666786cd721", registerDTO.Uuid)
}
