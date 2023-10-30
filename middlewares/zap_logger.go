package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func ZapLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		err := next(c)
		if err != nil {
			c.Error(err)
		}

		req := c.Request()
		res := c.Response()

		// Logging information
		remoteIP := c.RealIP()
		latency := time.Since(start)
		host := req.Host
		request := fmt.Sprintf("%s %s", req.Method, req.RequestURI)
		status := res.Status
		size := res.Size
		userAgent := req.UserAgent()

		id := req.Header.Get(echo.HeaderXRequestID)
		if id == "" {
			id = res.Header().Get(echo.HeaderXRequestID)
		}

		// Log messages based on status code
		switch {
		case status >= 500:
			log.Printf("Server error: %s - Remote IP: %s, Latency: %s, Host: %s, Request: %s, Status: %d, Size: %d, User Agent: %s, Request ID: %s",
				err.Error(), remoteIP, latency, host, request, status, size, userAgent, id)
		case status >= 400:
			log.Printf("Client error: %s - Remote IP: %s, Latency: %s, Host: %s, Request: %s, Status: %d, Size: %d, User Agent: %s, Request ID: %s",
				err.Error(), remoteIP, latency, host, request, status, size, userAgent, id)
		case status >= 300:
			log.Printf("Redirection - Remote IP: %s, Latency: %s, Host: %s, Request: %s, Status: %d, Size: %d, User Agent: %s, Request ID: %s",
				remoteIP, latency, host, request, status, size, userAgent, id)
		default:
			log.Printf("Success - Remote IP: %s, Latency: %s, Host: %s, Request: %s, Status: %d, Size: %d, User Agent: %s, Request ID: %s",
				remoteIP, latency, host, request, status, size, userAgent, id)
		}

		return nil
	}
}
