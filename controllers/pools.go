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

type MatchData struct {
    match services.Match
    picks []map[primitive.ObjectID]string
}

func (c *PoolsController) Get() {

    poolerId := c.GetSession("poolerId").(primitive.ObjectID)
    poolId := c.GetSession("poolId").(primitive.ObjectID)

    matches := services.NflAPI.FetchMatches(2021, 9)
    for _, match := range matches {
        log.Printf("[%v] > %v\n", match.EventId, match)
    }
    season, week := services.DB.FetchCurrentWeek(poolerId)
    poolers := services.DB.FetchPoolersFromPool(poolId)
    for i, p := range poolers {
        log.Printf("[%v] > %v\n", i, p)
    }

    poolPicks := services.DB.FetchPicks(poolers, season, week)
    for k, v := range poolPicks {
        log.Printf("[%v] > ...\n", k)
        for mid, pick := range v {
            log.Printf("\t[%v]: %v\n", mid, pick)
        }
    }

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = season
    c.Data["week"] = week

    c.Data["user"] = (c.GetSession("userId").(primitive.ObjectID)).Hex()
    c.Data["pooler"] = poolerId.Hex()
    //TODO: use MatchData struct to send this info in an object instead, easier for rendering
    c.Data["matches"] = matches
    //c.Data["picks"] = poolPicks

    c.Render()
}

