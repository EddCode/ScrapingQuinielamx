package main

import "ScrapingQuinielamx/calendarGames"

func main() {
	results := &calendarGames.ResultGames{}
	calendarGames.GamesResults(results)
	calendarGames.ShowGames(results)
}
