package dbprovider

import (
	"fmt"
	"time"

	"github.com/xdevices/register/dto"

	"github.com/xdevices/utilities/symbols"

	"github.com/xdevices/register/model"

	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	"github.com/xdevices/register/config"
	"github.com/xdevices/utilities/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Mgr DBManager

type DBManager interface {
	GetDb() *gorm.DB

	// Attributes
	GetAttributes() ([]model.Attribute, error)
	UpdateAttribute(attribute dto.AttributeDTO) (*model.Attribute, error)
	GetAttributeBySymbol(symbol string) (*model.Attribute, error)

	// Mappers
	MapAttributeToDTO(attribute *model.Attribute) dto.AttributeDTO
	MapToSensorAttribute(attributesDTO []dto.SensorAttributeDTO, now time.Time) []model.SensorAttribute
	MapToSensorAttributeDTO(entity *model.SensorAttribute) dto.SensorAttributeDTO
	MapToSensorEntity(registerDTO dto.SensorRegisterDTO) *model.SensorRegister
	MapToSensorDTO(sensor *model.SensorRegister) dto.SensorRegisterDTO

	// Sensor
	SaveSensor(sensor dto.SensorRegisterDTO) (*model.SensorRegister, error)
	GetAllSensors() ([]model.SensorRegister, error)
	GetSensorByUuid(uuid string) (*model.SensorRegister, error)
	UpdateSensor(sensorDTO dto.SensorRegisterDTO) (*model.SensorRegister, error)
	DeleteSensor(uuid string) (uint, error)
}

type manager struct {
	db *gorm.DB
}

func InitDbManager() {
	dbPath := config.Config().DBPath()

	if path, exists := db.AdjustDBPath(dbPath); !exists {
		dbPath = path
	}

	db2, err := gorm.Open("sqlite3", dbPath)

	if err != nil {
		errorMsg := fmt.Sprintf("Failed to init db[%s]", dbPath)
		log.Fatal(errorMsg, err)
	}
	db2.SingularTable(true)
	db2.AutoMigrate(&model.Attribute{}, &model.SensorAttribute{}, &model.SensorRegister{})
	initAttributes(db2)
	Mgr = &manager{db: db2}
}

func initAttributes(db *gorm.DB) {
	attribute := model.Attribute{}
	db.Where("symbol = ?", symbols.AcceptableMax.String()).First(&attribute)
	if attribute.Symbol == nil {
		insertAttribute(db, symbols.AcceptableMax.String(),
			"Acceptable maximum value",
			"Maximum value acceptable, before notification happens",
			"numeric")
	}

	attribute = model.Attribute{}
	db.Where("symbol = ?", symbols.AcceptableMin.String()).First(&attribute)
	if attribute.Symbol == nil {
		insertAttribute(db, symbols.AcceptableMin.String(),
			"Acceptable minimum value",
			"Minimum value acceptable, before notification happens",
			"numeric")
	}

	attribute = model.Attribute{}
	db.Where("symbol = ?", symbols.Active.String()).First(&attribute)
	if attribute.Symbol == nil {
		insertAttribute(db, symbols.Active.String(),
			"Active flag",
			"Is sensor active, or should it be ignored?",
			"boolean")
	}
}

func insertAttribute(db *gorm.DB, symbol, name, description, inputtype string) {
	db.Exec("insert into attribute (symbol, name, description, inputtype) values (?, ?, ?, ?)",
		symbol, name, description, inputtype)
}
