package main

import (
	"fmt"
	"github.com/labstack/echo"
	"helloworld/config"
	"net/http"
)

func main() {
	postgresConfig := config.NewPostgresConfig()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Hello,  %s", postgresConfig.DbName))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
