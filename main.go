package main

import (
	"CodeColony/internal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

// we got a map that is build as a coordinate system.

// we got a main building that can produce creeps

// we also got creeps that can move, collect and carry
// there should also be natural structures that the creep has to move around
// also there should definitly be more than one creep

func main() {
	go ech()
	go runGame()
	select {}
}
func runGame() {
	g := internal.NewGame(40, 30)
	g.Tick()
}

func ech() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
