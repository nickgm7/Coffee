package main

import (
	"github.com/labstack/echo/v4"
	"github.com/nickgm7/Coffee/handler"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHanlder{}
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3330")

}
