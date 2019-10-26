package sensor

import (
	"net/http"

	"github.com/maxzurawski/register/publishers"

	"github.com/maxzurawski/utilities/resterror"
	"github.com/maxzurawski/utilities/stringutils"

	"github.com/labstack/echo"
	"github.com/maxzurawski/register/services/sensor"
)

func HandleSensorDelete(c echo.Context) error {
	uuid := c.Param("uuid")

	if stringutils.IsZero(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("uuid parameter cannot be empty"))
	}

	if !stringutils.IsUuidValid(uuid) {
		return c.JSON(http.StatusBadRequest, resterror.New("invalid uuid given"))
	}

	sensorToDelete := sensor.Service.FindSensorByUuid(uuid)
	amount, err := sensor.Service.Delete(uuid)
	if err != nil && err.Error() != "record not found" {
		return c.JSON(http.StatusInternalServerError, resterror.New(err.Error()))
	}

	if err != nil && err.Error() == "record not found" {
		return c.NoContent(http.StatusNotFound)
	}

	if amount == 1 {
		publishers.SensorsPublisher().PublishDeleteChange(sensorToDelete, "")
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusNotFound)
}
