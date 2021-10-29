package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

    poolerId := c.GetSession("poolerId")

    if poolerId == nil {
        // TEMP EZ DEBUG
        //TODO: REMOVE THIS
        c.Redirect("/pools", 302)
    } else {
        c.Redirect("/pools", 302)
    }
}

