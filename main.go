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

	// Define route
	e.POST("/v1.0/transfer-va/create-va", handler.CreateVA)

	// Define route
	e.DELETE("/v1.0/transfer-va/delete-va", handler.DeleteVA)

	// Start server
	e.Logger.Fatal(e.Start(":7000"))
}
