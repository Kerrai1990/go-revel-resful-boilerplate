package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	greeting := "Hello world"

	return c.Render(greeting)
}

func (c App) Subscribe(email string) revel.Result {
	return c.Redirect(App.Index)
}
