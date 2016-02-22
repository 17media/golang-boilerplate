package main

import (
	"github.com/codegangsta/negroni"
	"github.com/dimfeld/httptreemux.git"
	"net/http"
)

// Handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	router = httptreemux.New()
	api := router.NewGroup("/api/v1")
	api.GET("/foo", hello) // becomes /api/v1/foo

	// Middleware
	n := negroni.Classic()
	n.UseHandler(router)

	// Start server
	n.Run(":3000")
}
