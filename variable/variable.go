package main

import "fmt"

func main() {
	var clubName string
	clubName = "Manchester United"
	fmt.Println(clubName)

	var established int
	established = 1878
	fmt.Println(established)

	var player = "Cristiano Ronaldo"
	var jerseyNumber = 7
	fmt.Println(player, jerseyNumber)

	// shorthand variable
	country := "Portugal"
	fmt.Println(country)

	// declare multi variables
	var (
		goals  = 100
		assist = 100
	)
	fmt.Println(goals, assist)
}
