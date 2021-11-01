package services

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type SportsDbAPI struct {
}

var NflAPI SportsDbAPI

var shortname_map = map[string]string{
    "ARI" : "Arizona Cardinals",
    "ATL" : "Atlanta Falcons",
    "BAL" : "Baltimore Ravens",
    "BUF" : "Buffalo Bills",
    "CAR" : "Carolina Panthers",
    "CHI" : "Chicago Bears",
    "CIN" : "Cincinnati Bengals",
    "CLE" : "Cleveland Browns",
    "DAL" : "Dallas Cowboys",
    "DEN" : "Denver Broncos",
    "DET" : "Detroit Lions",
    "GB" : "Green Bay Packers",
    "HOU" : "Houston Texans",
    "IND" : "Indianapolis Colts",
    "JAX" : "Jacksonville Jaguars",
    "KC" : "Kansas City Chiefs",
    "LA" : "Los Angeles Rams",
    "LAC" : "Los Angeles Chargers",
    "LV" : "Las Vegas Raiders",
    //"OAK" : "Oakland Raiders",
    "MIA" : "Miami Dolphins",
    "MIN" : "Minnesota Vikings",
    "NE" : "New England Patriots",
    "NO" : "New Orleans Saints",
    "NYG" : "New York Giants",
    "NYJ" : "New York Jets",
    "PHI" : "Philadelphia Eagles",
    "PIT" : "Pittsburgh Steelers",
    "SEA" : "Seattle Seahawks",
    "SF" : "San Francisco 49ers",
    "TB" : "Tampa Bay Buccaneers",
    "TEN" : "Tennessee Titans",
    "WAS" : "Washington",
}

var longname_map = map[string]string{
    "Arizona Cardinals" : "ARI",
    "Atlanta Falcons" : "ATL",
    "Baltimore Ravens" : "BAL",
    "Buffalo Bills" : "BUF",
    "Carolina Panthers" : "CAR",
    "Chicago Bears" : "CHI",
    "Cincinnati Bengals" : "CIN",
    "Cleveland Browns" : "CLE",
    "Dallas Cowboys" : "DAL",
    "Denver Broncos" : "DEN",
    "Detroit Lions" : "DET",
    "Green Bay Packers" : "GB",
    "Houston Texans" : "HOU",
    "Indianapolis Colts" : "IND",
    "Jacksonville Jaguars" : "JAX",
    "Kansas City Chiefs" : "KC",
    "Los Angeles Rams" : "LA",
    "Los Angeles Chargers" : "LAC",
    "Las Vegas Raiders" : "LV",
    "Oakland Raiders" : "LV",
    "Miami Dolphins" : "MIA",
    "Minnesota Vikings" : "MIN",
    "New England Patriots" : "NE",
    "New Orleans Saints" : "NO",
    "New York Giants" : "NYG",
    "New York Jets" : "NYJ",
    "Philadelphia Eagles" : "PHI",
    "Pittsburgh Steelers" : "PIT",
    "Seattle Seahawks" : "SEA",
    "San Francisco 49ers" : "SF",
    "Tampa Bay Buccaneers" : "TB",
    "Tennessee Titans" : "TEN",
    "Washington" : "WAS",
    "Washington Redskins" : "WAS",
}

func (s* SportsDbAPI) GetTeams() []string {
    i := 0
    teams := make([]string, len(shortname_map))

    for k := range shortname_map {
        teams[i] = k
        i++
    }

    return teams
}

func (s* SportsDbAPI) GetShortname(longname string) string {
    return longname_map[longname]
}

func (s* SportsDbAPI) GetLonganme(shortname string) string {
    return shortname_map[shortname]
}

func (s* SportsDbAPI) GetWeekName(weekNum int) string {
    if weekNum <= 18 {

        return fmt.Sprint(weekNum)
    } else if weekNum == 19 {

        return "WC"
    } else if weekNum == 20 {

        return "DV"
    } else if weekNum == 21 {

        return "CF"
    } else if weekNum == 22 {

        return "SB"
    }

    return "INVALID"
}

func (s* SportsDbAPI) GetWeekLongName(weekNum int) string {
    if weekNum <= 18 {

        return fmt.Sprintf("semaine %d", weekNum)
    } else if weekNum == 19 {

        return "WildCards"
    } else if weekNum == 20 {

        return "Division Round"
    } else if weekNum == 21 {

        return "Conference Championship"
    } else if weekNum == 22 {

        return "SuperBowl"
    }

    return "INVALID"
}

//--------------------------------------------------

type ParsedMatch struct {
    EventId string `json:"idEvent"`
    EventName string `json:"strEvent"`
    EventNameAlt string `json:"strEventAlternate"`

    Season string `json:"strSeason"`
    Week string `json:"intRound"`

    HomeTeam string `json:"strHomeTeam"`
    HomeScore string `json:"intHomeScore"`
    AwayTeam string `json:"strAwayTeam"`
    AwayScore string `json:"intAwayScore"`
}

type Match struct {
    EventId string
    EventName string
    EventNameAlt string

    Season int
    Week int

    HomeTeam string
    HomeScore int
    AwayTeam string
    AwayScore int
}

func (parsed *ParsedMatch) convert() Match {
    var m Match

    m.EventId = parsed.EventId
    m.EventName = parsed.EventName
    m.EventNameAlt = parsed.EventNameAlt

    season, _ := strconv.Atoi(parsed.Season)
    m.Season = season
    week, _ := strconv.Atoi(parsed.Week)
    m.Week = week

    m.HomeTeam = parsed.HomeTeam
    m.AwayTeam = parsed.AwayTeam

    homeScore, _ := strconv.Atoi(parsed.HomeScore)
    m.HomeScore = homeScore
    awayScore, _ := strconv.Atoi(parsed.AwayScore)
    m.AwayScore = awayScore

    return m
}

func (s* SportsDbAPI) FetchMatches(season int, week int) []Match {
    dataUrl, _ := beego.AppConfig.String("THESPORTDB_URL")
    url := fmt.Sprintf(dataUrl, season, week)

    res, err := http.Get(url)
    if err != nil {
        log.Fatalf("[TheSportsDB] Could not get url: %v\nerror: %v\n", url, err)
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)

    data := make(map[string][]ParsedMatch)
    json.Unmarshal(body, &data)

    results := make([]Match, 0, 16)
    for _, p := range data["events"] {
        results = append(results, p.convert())
    }

    return results
}
