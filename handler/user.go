package handler

import (
	"github.com/labstack/echo/v4"
)

type UserHanlder struct{}

func (h UserHanlder) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
