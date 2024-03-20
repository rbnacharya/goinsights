package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/rbnacharya/trafficinsights-go/internal/controller"
	"github.com/rbnacharya/trafficinsights-go/internal/core"

	_ "github.com/lib/pq"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	err := v.validator.Struct(i)

	if err == nil {
		return nil
	}

	return err.(validator.ValidationErrors)
}

func main() {

	e := echo.New()

	// Add validator to app context
	e.Validator = &Validator{validator: validator.New()}

	err, db := core.InitDb()

	if err != nil {
		e.Logger.Fatal(err)
	}

	global := core.NewGlobal(db)

	ctrl := &controller.Controllers{
		Global: global,
	}

	e.POST("/request", ctrl.RegisterClick)
	e.GET("/statistics", ctrl.FetchStats)

	e.Logger.Fatal(e.Start(":8080"))
}
