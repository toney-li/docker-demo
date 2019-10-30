package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", handlerindex)
	fmt.Println("starting echo")
	err := e.Start(":8080")
	if err != nil {
		fmt.Println("echo", err)
	}
}

func handlerindex(c echo.Context) error {
	fmt.Println("Hello world handlerindex")
	return c.JSON(http.StatusOK, `{"hello":"world"}`)
}
