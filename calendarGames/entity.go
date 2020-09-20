package calendarGames

import "fmt"

type ResultGames struct {
	games []*game
}

type game struct {
	HomeTeam  string
	VisitTeam string
	Result    string
	Journey   int16
}

type CalendarGame interface {
	add()
	listGames()
}

func (this *ResultGames) add(gameResult *game) {
	this.games = append(this.games, gameResult)
}

func (this *ResultGames) listGames() {
	for idx, game := range this.games {
		fmt.Printf("Game %d Home %s %s  Visit %s \n", idx+1, game.HomeTeam, game.Result, game.VisitTeam)
	}
}
