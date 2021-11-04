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

    matches := services.NflAPI.FetchMatches(2021, 1)
    for i, m := range matches {
        log.Printf("[%v] %v\n", i, m)
    }

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = 9999
    c.Data["week"] = 99

    c.Data["user"] = c.GetSession("userId")
    c.Data["pooler"] = c.GetSession("poolerId")
    c.Data["matches"] = matches

    c.Render()
}

