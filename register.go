package main

import (
	"github.com/labstack/echo"
	"github.com/xdevices/register/config"
	"github.com/xdevices/register/dbprovider"
	attribute2 "github.com/xdevices/register/handlers/attribute"
	sensor2 "github.com/xdevices/register/handlers/sensor"
	"github.com/xdevices/register/publishers"
	"github.com/xdevices/register/services/attribute"
	"github.com/xdevices/register/services/sensor"
)

func init() {
	eureka := config.EurekaManagerInit()
	eureka.SendRegistrationOrFail()
	eureka.ScheduleHeartBeat(config.Config().ServiceName(), 10)

	dbprovider.InitDbManager()
	attribute.Init()
	sensor.Init()
	publishers.Init()
}

func main() {

	e := echo.New()

	// Sensors
	e.GET("/sensors/", sensor2.HandleGetSensors)
	e.POST("/sensors/", sensor2.HandleSaveSensor)
	e.GET("/sensors/:uuid", sensor2.HandleGetSensorByUuid)
	e.PUT("/sensors/:uuid", sensor2.HandleSensorUpdate)
	e.DELETE("/sensors/:uuid", sensor2.HandleSensorDelete)

	// Attributes
	e.GET("/attributes/", attribute2.HandleGetAllAttributes)
	e.GET("/attributes/:symbol", attribute2.HandleGetAttributeBySymbol)
	e.PUT("/attributes/:symbol", attribute2.HandleUpdateAttribute)

	e.Logger.Fatal(e.Start(config.Config().Address()))
}
