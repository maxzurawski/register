package attribute

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xdevices/register/services/attribute"
)

func HandleGetAllAttributes(c echo.Context) error {

	dtos, err := attribute.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(dtos) == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}

	return c.JSON(http.StatusOK, dtos)
}
