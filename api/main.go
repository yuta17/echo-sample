package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	subscribe struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}
)

var (
	subscribes = map[int]*subscribe{}
	seq        = 1
)

//----------
// Handlers
//----------

func createSubscribe(c echo.Context) error {
	s := &subscribe{
		ID: seq,
	}
	if err := c.Bind(s); err != nil {
		return err
	}
	subscribes[s.ID] = s
	seq++
	return c.JSON(http.StatusCreated, s)
}

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.POST("/subscribes", createSubscribe)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
