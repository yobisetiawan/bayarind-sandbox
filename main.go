package main

import (
	"bayarind-sandbox/handler"

	appMiddleware "bayarind-sandbox/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(appMiddleware.RequestLoggerMiddleware)
	e.Use(middleware.Logger())

	e.POST("/v1.0/transfer-va/create-va", handler.CreateVA)

	e.DELETE("/v1.0/transfer-va/delete-va", handler.DeleteVA)

	e.POST("/v1.0/transfer-va/status", handler.StatusVa)

	// Start server
	e.Logger.Fatal(e.Start(":7000"))
}
