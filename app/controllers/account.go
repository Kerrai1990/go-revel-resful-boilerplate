package controllers

import (
	"github.com/revel/revel"
)

type Account struct {
	*revel.Controller
}

func (c Account) Index() revel.Result {
	return c.Render()
}
