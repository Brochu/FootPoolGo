package controllers

import (
    "log"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    c.Data["test0"] = c.GetSession("oauthState")
    c.Data["test1"] = c.GetSession("userId")
    c.Data["test2"] = c.GetSession("poolerId")

    res := services.FetchMathes(2021, 1)
    for _, m := range res {
        log.Printf("* %v\n", m)
    }

    teams := services.GetTeams()
    for _, t := range teams {
        log.Printf("* %v\n", t)
    }

    c.TplName = "listPoolsTemplate.tpl"
    c.Layout = "layout.html"

    c.Render();
}

