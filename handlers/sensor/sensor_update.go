package sensor

import (
	"net/http"
	"strings"

	"github.com/xdevices/register/publishers"

	"github.com/xdevices/utilities/resterror"
	"github.com/xdevices/utilities/stringutils"

	"github.com/labstack/echo"
	"github.com/xdevices/register/dto"
	"github.com/xdevices/register/services/sensor"
)

func HandleSensorUpdate(c echo.Context) error {

	uuid := c.Param("uuid")

	if stringutils.IsZero(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("uuid parameter cannot be empty"))
	}

	if !stringutils.IsUuidValid(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("invalid uuid given"))
	}

	sensorDTO := &dto.SensorRegisterDTO{}
	if err := c.Bind(sensorDTO); err != nil {
		if strings.Contains(err.Error(), "415") {
			return c.JSON(http.StatusUnsupportedMediaType, resterror.New(err.Error()))
		}
		return c.JSON(http.StatusBadRequest, resterror.New(err.Error()))
	}

	oldSensor := sensor.Service.FindSensorByUuid(uuid)
	if oldSensor.Version != sensorDTO.Version {
		return c.JSON(http.StatusBadRequest, resterror.New("optimistic lock exception. versions are different"))
	}

	updated, err := sensor.Service.Update(*sensorDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resterror.New(err.Error()))
	}

	publishers.SensorsPublisher().PublishUpdateChange(oldSensor, updated)

	return c.JSON(http.StatusAccepted, updated)

}
