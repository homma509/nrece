package handler

import (
	"fmt"

	"github.com/homma509/nrece/infra"
	"github.com/homma509/nrece/usecase"
	"github.com/labstack/echo"
)

type apiError struct {
	Code    int
	Message string
}

// AppHandler ...
type AppHandler struct {
	*usecase.App
}

// NewAppHandler ...
func NewAppHandler() *AppHandler {
	return &AppHandler{
		usecase.NewApp(
			infra.NewAppRepository(
				infra.NewSQLHandler(),
			),
		),
	}
}

// GetApp ...
func (handler *AppHandler) GetApp() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		id := c.QueryParam("id")
		if id == "" {
			return c.JSON(400, "Invalid parameter id.")
		}

		//err = utils.HeaderCheck(c, echo.HeaderContentType, echo.MIMEApplicationJSON)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		//err = utils.ClientIDCheck(c)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		resJSON, err := handler.Get(id)
		if err != nil {
			return c.JSON(503,
				apiError{
					Code:    503,
					Message: fmt.Sprintf("AppHandler failed to GetApp %v", err),
				},
			)
			// return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(200, resJSON)
	}
}
