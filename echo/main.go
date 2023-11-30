package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/orders", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "MY HOMEPAGE. SEE <a href=\"/orders\">ORDERS</a>")
	})

	e.GET("/orders", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &InternalStorageKeys())
	})

	e.GET("/orders/:id", func(c echo.Context) error {
		// order ID from path `order/:id`
		id := c.Param("id")
		return c.JSON(http.StatusOK, InternalStorageGet(id))
	})
	e.Logger.Fatal(e.Start(":8880"))
}
