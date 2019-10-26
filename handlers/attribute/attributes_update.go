package attribute

import (
	"net/http"
	"strings"

	"github.com/maxzurawski/register/publishers"

	"github.com/maxzurawski/utilities/resterror"

	"github.com/labstack/echo"
	"github.com/maxzurawski/register/dto"
	"github.com/maxzurawski/register/services/attribute"
	"github.com/maxzurawski/utilities/stringutils"
)

func HandleUpdateAttribute(c echo.Context) error {
	symbol := c.Param("symbol")

	if stringutils.IsZero(symbol) {
		return c.JSON(http.StatusBadRequest, resterror.New("not known symbol"))
	}

	attributeDTO := &dto.AttributeDTO{}
	if err := c.Bind(attributeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, resterror.New(err.Error()))
	}

	if stringutils.IsZero(attributeDTO.Symbol) {
		return c.JSON(http.StatusBadRequest, resterror.New("given symbol of attribute cannot be zero or empty"))
	}

	if strings.ToUpper(attributeDTO.Symbol) != strings.ToUpper(symbol) {
		return c.JSON(http.StatusBadRequest, resterror.New("given symbol with symbol in attribute provided does not match"))
	}

	oldAttribute, _ := attribute.Service.GetAttributeBySymbol(symbol)
	updateAttribute, err := attribute.Service.UpdateAttribute(*attributeDTO)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resterror.New(err.Error()))
	}

	publishers.AttributesPublisher().PublishUpdateChange(oldAttribute, updateAttribute)

	return c.JSON(http.StatusAccepted, updateAttribute)
}
