package main

import (
	"net/http"
	"os"
	"github.com/labstack/echo"
	rice "github.com/GeertJohan/go.rice"
)

func main() {
	e := echo.New()
	env := os.Getenv("ENV")

	if env == "production" {
		assetHandler := http.FileServer(rice.MustFindBox("web/build").HTTPBox())

		e.GET("/", echo.WrapHandler(assetHandler))

		e.GET("/static/*", echo.WrapHandler(assetHandler))

	}

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
