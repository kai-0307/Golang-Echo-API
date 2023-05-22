package main

import (
	"go-echo-api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go!")
	})

	u := user.New(1, "golangUser", 20)

	e.GET("/user", u.Get)
	e.POST("/user", user.Post)

	e.Logger.Fatal(e.Start(":1234"))
}
