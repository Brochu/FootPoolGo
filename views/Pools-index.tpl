<main role="main" class="container">

  <div class="starter-template">
    <h2>Pour saison {{.season}} - semaine {{.week}}</h2>

    <p class="lead">LISTING POOLS HERE</p>
    <!--
    backed by this struct:
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
    -->
    <div class="match-list">
        {{ range .pooldata }}
            <div>{{.Match.AwayTeam}}</div>
            <div>{{.Match.AwayScore}}</div>

            <div>{{.Match.HomeScore}}</div>
            <div>{{.Match.HomeTeam}}</div>

            <div>{{.Match.EventId}}</div>
            <div>
                {{ range $k, $v := .Picks }}
                    <p>{{$v}}</p>
                {{ end }}
            </div>
        {{ end }}
    </div>

    <p>for user/pooler: </p>
    <p>{{.user}} / {{.pooler}}</p>
  </div>

</main>

