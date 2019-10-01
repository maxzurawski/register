package sensor

import (
	"net/http"

	"github.com/xdevices/utilities/resterror"

	"github.com/labstack/echo"
	"github.com/xdevices/register/services/sensor"
)

func HandleGetSensors(c echo.Context) error {

	dtos, err := sensor.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resterror.ErrorMsg{Msg: err.Error()})
	}
	if len(dtos) == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}
	return c.JSON(http.StatusOK, dtos)
}
