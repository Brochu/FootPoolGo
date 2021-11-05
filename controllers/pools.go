package controllers

import (
    "log"
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    poolerId := c.GetSession("poolerId").(primitive.ObjectID)

    matches := services.NflAPI.FetchMatches(2021, 1)
    for i, m := range matches {
        log.Printf("[%v] %v\n", i, m)
    }

    picks := services.DB.FetchAllPicksCurrentWeek(poolerId)
    for i, p := range picks {
        log.Printf("[%v] %v", i, p)
    }

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = 9999
    c.Data["week"] = 99

    c.Data["user"] = (c.GetSession("userId").(primitive.ObjectID)).Hex()
    c.Data["pooler"] = poolerId.Hex()
    c.Data["matches"] = matches

    c.Render()
}

