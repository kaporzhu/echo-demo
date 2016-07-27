package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"

	demo "github.com/kaporzhu/echo-demo"
	controller "github.com/kaporzhu/echo-demo/controller"
)

func main() {
	e := demo.New()

	// Middleware
	e.Use(middleware.Logger(), middleware.Recover(), middleware.AddTrailingSlash())

	// Routes
	controller.SampleController.Load(e)

	// Start server
	e.Run(standard.New(":1323"))
}
