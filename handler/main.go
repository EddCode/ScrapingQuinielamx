package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

func Root(w http.ResponseWriter, request *http.Request) {
	gamesResults()
	io.WriteString(w, "Web scrapint starting")
}

func gamesResults() {
	url := "https://www.mediotiempo.com/futbol/liga-mx/calendario"
	fmt.Print(url)

	defer func() {
		recov := recover()
		log.Println(recov)
	}()

	if url == "" {
		panic("Missing URL argument to starting scraping")
	}

	log.Println("Reading website => ", url)
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

	collector.Visit(url)

	_, err := json.Marshal(scrapData)

	if err != nil {
		panic(err)
	}

}
