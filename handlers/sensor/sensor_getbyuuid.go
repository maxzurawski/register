package sensor

import (
	"net/http"

	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"
	"github.com/xdevices/register/services/sensor"
	"github.com/xdevices/utilities/stringutils"
)

func HandleGetSensorByUuid(c echo.Context) error {
	uuid := c.Param("uuid")
	if stringutils.IsZero(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("uuid parameter cannot be empty"))
	}

	if !stringutils.IsUuidValid(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("invalid uuid given"))
	}

	sensor := sensor.Service.FindSensorByUuid(uuid)
	if sensor == nil {
		return c.JSON(http.StatusNoContent, resterror.New("unknown uuid given"))
	}
	return c.JSON(http.StatusOK, sensor)
}
