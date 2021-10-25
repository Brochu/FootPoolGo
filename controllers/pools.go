package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    c.Data["test"] = c.GetSession("oauthState")
    c.TplName = "listPoolsTemplate.tpl"
    c.Layout = "layout.html"

    c.Render();
}

