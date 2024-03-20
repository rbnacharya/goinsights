package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleError(boom *Booms, ec echo.Context) {
	ec.JSON(http.StatusBadRequest, boom)
}
