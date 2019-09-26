package dbprovider

import (
	"strings"
	"testing"
	"time"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/assert"

	"github.com/xdevices/utilities/symbols"

	"github.com/stretchr/testify/suite"
	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/model"
)

// suite struct
type SaveSensorSuite struct {
	suite.Suite
}

// init suite
func TestSaveSensorSuite(t *testing.T) {
	suite.Run(t, new(SaveSensorSuite))
}

// setup test function
func (s *SaveSensorSuite) SetupTest() {
	EnvironmentPreparations()
}

// prepare method for creating sensor register with sensor attributes
func createSensorRegister(id uint, registerUuid, name, description, sensortype string) dto.SensorRegisterDTO {

	version := uint(0)
	stUpper := strings.ToUpper(sensortype)

	register := model.SensorRegister{
		ID:          &id,
		Version:     &version,
		Name:        &name,
		Description: &description,
		Type:        &stUpper,
		Uuid:        &registerUuid,
		Attributes:  []model.SensorAttribute{},
	}
	Mgr.GetDb().Save(register)
	Mgr.GetDb().Where("uuid=?", registerUuid).Find(&register)

	var attributesDTO []dto.SensorAttributeDTO

	attribute := dto.SensorAttributeDTO{}
	attribute.Value = "true"
	attribute.SensorRegisterID = *register.ID
	attribute.Version = 0
	attribute.Symbol = symbols.Active.String()
	attributesDTO = append(attributesDTO, attribute)

	attribute2 := dto.SensorAttributeDTO{}
	attribute2.Value = "24.5"
	attribute2.Symbol = symbols.AcceptableMin.String()
	attribute2.Version = 0
	attribute2.SensorRegisterID = *register.ID
	attributesDTO = append(attributesDTO, attribute2)

	attributes := Mgr.MapToSensorAttribute(attributesDTO, time.Now())
	register.Attributes = attributes
	Mgr.GetDb().Save(register)

	return Mgr.MapToSensorDTO(&register)
}

// prepare method for creating test sensors - as if sensors would already exist
func prepareFakeSensorRegisters() (dtos []dto.SensorRegisterDTO) {

	dto := createSensorRegister(1, "b255382c-35c8-4aee-99db-74fbe41ca9f7", "test 1", "test 1 description", "dummy_type")
	dtos = append(dtos, dto)

	dto = createSensorRegister(2, "a607f675-e106-4617-9261-bcd5ee5f2ad7", "test 2", "test 2 description", "test2_type")
	dtos = append(dtos, dto)

	return
}

// failure test, when given sensor register id is not 0
func (s *SaveSensorSuite) TestFailure_RegisterDTO_hasID() {

	// Arrange.

	dto := dto.SensorRegisterDTO{
		ID:   12,
		Name: "Testing failure",
	}

	// Act.

	result, err := Mgr.SaveSensor(dto)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), result)
	assert.Equal(s.T(), "sensor with given id exists. no save. try to update", err.Error())
}

// failure test, because uuid is invalid
func (s *SaveSensorSuite) TestFailure_Uuid_NotValid() {

	// Arrange.

	dto := dto.SensorRegisterDTO{
		ID:   0,
		Name: "Testing failure",
		Uuid: "not_valid_uuid",
	}

	// Act.

	result, err := Mgr.SaveSensor(dto)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), result)
	assert.Equal(s.T(), "uuid is not valid", err.Error())
}

// failure test, because sensorRegister.ID already exists (REMEMBER to use cleaner hook, when preparing test sensors (!!!))
func (s *SaveSensorSuite) TestFailure_RegisterExists() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	fakeRegisters := prepareFakeSensorRegisters()
	dtoToSave := fakeRegisters[0]
	dtoToSave.Uuid = "b255382c-35c8-4aee-99db-74fbe41ca9f7"
	dtoToSave.ID = 0

	// Act.

	result, err := Mgr.SaveSensor(dtoToSave)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), result)
	assert.Equal(s.T(), "ups, something went wrong. sensor register with uuid=[b255382c-35c8-4aee-99db-74fbe41ca9f7] already exists", err.Error())
}

// success test (REMEMBER to use cleaner hook, when preparing test sensors (!!!))
func (s *SaveSensorSuite) TestSuccess() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	fakeRegisters := prepareFakeSensorRegisters()
	dtoToSave := fakeRegisters[0]
	dtoToSave.ID = 0
	dtoToSave.Type = "dummy_type"
	dtoToSave.Uuid = "9c73f3c1-0830-4a5f-9baf-d904f0e37536"

	// Act.

	result, err := Mgr.SaveSensor(dtoToSave)

	// Assert.

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.NotNil(s.T(), result.ID)
	assert.True(s.T(), *result.ID > 0)
	assert.Equal(s.T(), "9c73f3c1-0830-4a5f-9baf-d904f0e37536", *result.Uuid)
}
