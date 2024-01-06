package handler

import (
	"github.com/anousonefs/templ-basic/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}
