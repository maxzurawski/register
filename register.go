package main

import (
	"github.com/labstack/echo"
	"github.com/maxzurawski/register/config"
	"github.com/maxzurawski/register/dbprovider"
	attribute2 "github.com/maxzurawski/register/handlers/attribute"
	sensor2 "github.com/maxzurawski/register/handlers/sensor"
	"github.com/maxzurawski/register/publishers"
	"github.com/maxzurawski/register/services/attribute"
	"github.com/maxzurawski/register/services/sensor"
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

	// cache
	e.GET("/cachesensors/", sensor2.HandleGetCacheSensors)

	e.Logger.Fatal(e.Start(config.Config().Address()))
}
