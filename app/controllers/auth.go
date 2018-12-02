package controllers

import (
	"github.com/revel/revel"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Login(email string, password string) revel.Result {
	return c.Render()
}

func (c Auth) Register(email string, password string) revel.Result {
	return c.Render()
}

func (c Auth) LoginPost(email string, password string) revel.Result {
	return c.Render()
}

func (c Auth) RegisterPost(email string, password string) revel.Result {
	return c.Redirect(App.Index)
}
