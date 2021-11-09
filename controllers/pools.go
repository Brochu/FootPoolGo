package controllers

import (
    "FootPoolGo/services"

	beego "github.com/beego/beego/v2/server/web"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolsController struct {
	beego.Controller
}

type MatchData struct {
    Match services.Match
    Picks map[primitive.ObjectID]string
}

func (c *PoolsController) Get() {

    poolerId := c.GetSession("poolerId").(primitive.ObjectID)
    poolId := c.GetSession("poolId").(primitive.ObjectID)

    matches := services.NflAPI.FetchMatches(2021, 9)
    season, week := services.DB.FetchCurrentWeek(poolerId)
    poolers := services.DB.FetchPoolersFromPool(poolId)

    poolPicks := services.DB.FetchPicks(poolers, season, week)

    poolData := make([]MatchData, len(matches))
    for i, match := range matches {
        poolData[i].Match = match
        poolData[i].Picks = poolPicks[match.EventId]
    }

    c.Layout = "layout.html"
    c.TplName = "Pools-index.tpl"

    c.Data["season"] = season
    c.Data["week"] = week

    c.Data["poolers"] = poolers;
    c.Data["pooldata"] = poolData

    c.Render()
}

