package main

import (
	"fmt"

	"github.com/anousonefs/templ-basic/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.GET("/user", userHandler.HandlerUserShow)

	fmt.Println("hello templ")

	app.Start(":8080")
}
