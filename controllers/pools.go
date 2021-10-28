package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    c.Data["test0"] = c.GetSession("oauthState")
    c.Data["test1"] = c.GetSession("userId")
    c.Data["test2"] = c.GetSession("poolerId")

    c.TplName = "listPoolsTemplate.tpl"
    c.Layout = "layout.html"

    c.Render();
}

