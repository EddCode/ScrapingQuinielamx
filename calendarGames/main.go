package calendarGames

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func GamesResults() {
	url := "https://www.mediotiempo.com/futbol/liga-mx/calendario"

	defer func() {
		recov := recover()
		log.Printf("Error: %s \n", recov)
	}()

	if url == "" {
		panic("Missing URL argument to starting scraping")
	}

	collector := colly.NewCollector()

	var scrapData []string
	collector.OnHTML(".going.teams ", func(readerHtml *colly.HTMLElement) {

		firstTeam := readerHtml.ChildText(".first-team > .team-name.large span")
		secondTeam := readerHtml.ChildText(".second-team > .team-name.large span")
		resultTeam := readerHtml.ChildText(".result-team span")

		fmt.Printf("Equipo local: %s Equipo Visitante: %s result => %s \n", firstTeam, secondTeam, resultTeam)

		if firstTeam != "" {
			scrapData = append(scrapData, firstTeam)
		}
	})

	// TODO: make struct with game data eg. {home team: lose, visit team: win, result, game status}
	// save each struct at slice
	collector.Visit(url)

	_, err := json.Marshal(scrapData)

	if err != nil {
		panic(err)
	}

}
