package controllers

import (
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolsController struct {
	beego.Controller
}

func (c *PoolsController) Get() {

    poolerId := c.GetSession("poolerId").(primitive.ObjectID)
    poolId := c.GetSession("poolId").(primitive.ObjectID)

    matches := services.NflAPI.FetchMatches(2021, 9)
    season, week := services.DB.FetchCurrentWeek(poolerId)

    poolPicks := services.DB.FetchPoolPicks(poolId, season, week)
    //for matchId, picksMap := range poolPicks {
    //    log.Printf("Picks for match %v: \n", matchId)

    //    for poolerId, pick := range picksMap {
    //        log.Printf("\t[%v] > %v\n", poolerId, pick)
    //    }
    //}

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = season
    c.Data["week"] = week

    c.Data["user"] = (c.GetSession("userId").(primitive.ObjectID)).Hex()
    c.Data["pooler"] = poolerId.Hex()
    c.Data["matches"] = matches
    c.Data["picks"] = poolPicks

    c.Render()
}

