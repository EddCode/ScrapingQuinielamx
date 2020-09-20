package calendarGames

import (
	"log"

	"github.com/gocolly/colly"
)

func GamesResults(calendarGame *ResultGames) {
	url := "https://www.mediotiempo.com/futbol/liga-mx/calendario"

	defer func() {
		recov := recover()

		if recov != nil {
			log.Printf("Error: %s \n", recov)
		}
	}()

	if url == "" {
		panic("Missing URL argument to starting scraping")
	}

	collector := colly.NewCollector()

	collector.OnHTML(".going.teams ", func(readerHtml *colly.HTMLElement) {

		firstTeam := readerHtml.ChildText(".first-team > .team-name.large span")
		secondTeam := readerHtml.ChildText(".second-team > .team-name.large span")
		resultTeam := readerHtml.ChildText(".result-team span")

		currentGame := &game{
			HomeTeam:  firstTeam,
			VisitTeam: secondTeam,
			Result:    resultTeam,
			Journey:   1,
		}

		calendarGame.add(currentGame)

	})

	collector.Visit(url)

}

func ShowGames(calendarGame *ResultGames) {
	calendarGame.listGames()
}
