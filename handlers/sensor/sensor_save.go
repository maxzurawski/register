package sensor

import (
	"net/http"

	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"
	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/services/sensor"
)

func HandleSaveSensor(c echo.Context) error {

	sensorDTO := &dto.SensorRegisterDTO{}
	if err := c.Bind(sensorDTO); err != nil {
		return c.JSON(http.StatusBadRequest, resterror.ErrorMsg{Msg: err.Error()})
	}
	created, err := sensor.Service.Save(*sensorDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resterror.ErrorMsg{Msg: err.Error()})
	}

	return c.JSON(http.StatusCreated, created)
}
