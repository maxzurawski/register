package sensor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxzurawski/utilities/symbols"

	"github.com/maxzurawski/utilities/db"

	"github.com/stretchr/testify/assert"

	"github.com/maxzurawski/utilities/resterror"

	"github.com/maxzurawski/register/dto"

	"github.com/labstack/echo"

	"github.com/maxzurawski/register/dbprovider"
	"github.com/maxzurawski/register/services/sensor"
	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleSensorSaveSuite struct {
	suite.Suite
}

// init suite
func TestHandleSensorSaveSuite(t *testing.T) {
	suite.Run(t, new(HandleSensorSaveSuite))
}

// setup test func
func (h *HandleSensorSaveSuite) SetupTest() {
	dbprovider.EnvironmentPreparations()
	sensor.Init()
}

// failure test - body incorrect
func (h *HandleSensorSaveSuite) TestFailure_BodyIncorrect() {

	// Arrange.

	var sensors []dto.SensorRegisterDTO
	sensors = append(sensors, dto.SensorRegisterDTO{})
	bytes, _ := json.Marshal(sensors)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleSaveSensor(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, rec.Code)
	var err resterror.ErrorMsg
	_ = json.NewDecoder(rec.Body).Decode(&err)
	assert.Equal(h.T(), "code=400, message=Unmarshal type error: expected=dto.SensorRegisterDTO, got=array, field=, offset=1", err.Msg)
}

// success test
func (h *HandleSensorSaveSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()

	var attributes []dto.SensorAttributeDTO
	attributes = append(attributes, dto.SensorAttributeDTO{Symbol: symbols.Active.String(), Value: "true"})

	sensor := dto.SensorRegisterDTO{
		Name:       "Dummy sensor",
		Type:       "DUMMY_TYPE",
		Uuid:       "81750491-88dd-410e-b53f-1666786cd721",
		Attributes: attributes,
	}
	bytes, _ := json.Marshal(sensor)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleSaveSensor(c)

	// Assert.

	assert.Equal(h.T(), http.StatusCreated, rec.Code)
	var afterSave dto.SensorRegisterDTO
	_ = json.NewDecoder(rec.Body).Decode(&afterSave)
	assert.NotNil(h.T(), afterSave)
	assert.NotNil(h.T(), afterSave.ID)
	assert.True(h.T(), afterSave.ID > 0)
	assert.True(h.T(), afterSave.Version == 0)
	assert.True(h.T(), len(afterSave.Attributes) == 1)
	attributeDTO := afterSave.Attributes[0]
	assert.True(h.T(), attributeDTO.ID > 0)
	assert.True(h.T(), attributeDTO.SensorRegisterID == afterSave.ID)

}
