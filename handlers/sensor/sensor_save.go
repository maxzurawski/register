package sensor

import (
	"net/http"

	"github.com/maxzurawski/register/publishers"

	"github.com/maxzurawski/utilities/resterror"

	"github.com/labstack/echo"
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/services/sensor"
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

	publishers.SensorsPublisher().PublishSaveChange("", created)

	return c.JSON(http.StatusCreated, created)
}
