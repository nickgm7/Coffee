package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/nickgm7/Coffee/view/user"
)

type UserHanlder struct{}

func (h UserHanlder) HandleUserShow(c echo.Context) error {
	return render(c, user.Show())
}
