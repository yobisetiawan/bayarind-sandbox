package main

import (
	"bayarind-sandbox/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Define route
	e.POST("/v1.0/transfer-va/create-va", handler.CreateVA)

	// Define route
	e.DELETE("/v1.0/transfer-va/delete-va", handler.DeleteVA)

	// Start server
	e.Logger.Fatal(e.Start(":7000"))
}
