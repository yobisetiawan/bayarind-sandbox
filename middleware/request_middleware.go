package appMiddleware

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get request details
		start := time.Now()
		req := c.Request()
		res := c.Response()

		// Log request details before processing
		log.Printf("Started %s %s from %s", req.Method, req.RequestURI, req.RemoteAddr)

		// Call the next handler in the chain
		err := next(c)

		// Log response details after processing
		log.Printf("Process %s %s from %s, Completed %d %s in %v", req.Method, req.RequestURI, req.RemoteAddr, res.Status,
			http.StatusText(res.Status),
			time.Since(start))

		return err
	}
}
