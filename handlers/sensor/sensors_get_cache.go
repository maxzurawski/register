package sensor

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/services/sensor"
)

func HandleGetCacheSensors(c echo.Context) error {
	dtos, err := sensor.Service.GetCachedSensors()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(dtos) == 0 {
		return c.JSON(http.StatusNoContent, []dto.CachedSensorDTO{})
	}
	return c.JSON(http.StatusOK, dtos)
}
