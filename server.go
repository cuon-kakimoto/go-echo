package main

import (
	"github.com/cuon-kakimoto/go-echo/hanlder"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var (
	mockDB = map[string]*User{
		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
	}
)

func main() {
	// Echo instance
	e := echo.New()
	h := handler.Handler{mockDB}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/", h.GetUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
