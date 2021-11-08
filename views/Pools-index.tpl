<main role="main" class="container">

    <div class="starter-template">
        <h2>Pour saison {{.season}} - semaine {{.week}}</h2>

        <div class="match-list">
            {{ range .pooldata }}
                <!-- MATCH INFO -->
                <div>{{.Match.AwayTeam}}</div>
                <div>{{.Match.AwayScore}}</div>

                <div>{{.Match.HomeScore}}</div>
                <div>{{.Match.HomeTeam}}</div>

                <!-- PICKS INFO -->
                <div class="picks-list">
                {{ range $k, $v := .Picks }}
                    <span>{{$v}}</span>
                {{ end }}
            </div>
            {{ end }}
        </div>
    </div>

</main>

