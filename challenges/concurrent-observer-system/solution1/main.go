package main

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.RequestLogger())

	eventManager := registerObservers()

	e.POST("/deployment", func(c *echo.Context) error {
		u := new(Event)
		if err := c.Bind(u); err != nil {
			return err
		}

		eventManager.notifyAll(*u)

		return c.JSON(http.StatusCreated, u)
	})

	if err := e.Start(":1323"); err != nil {
		e.Logger.Info("shutting down the server")
	}
}

func registerObservers() *EventManager {
	eventManager := &EventManager{}
	eventManager.register(&DeploymentLogger{})
	eventManager.register(&DeploymentAudit{})
	eventManager.register(&DeploymentNotification{})
	return eventManager
}
