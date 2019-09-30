package attribute

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xdevices/register/services/attribute"
)

func HandleGetAttributeBySymbol(c echo.Context) error {
	symbol := c.Param("symbol")
	dto, err := attribute.Service.GetAttributeBySymbol(symbol)

	if err != nil && err.Error() != "record not found" {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err != nil && err.Error() == "record not found" {
		return c.JSON(http.StatusBadRequest, nil)
	}

	return c.JSON(http.StatusOK, dto)
}
