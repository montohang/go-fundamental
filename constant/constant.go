package main

import "fmt"

func main() {
	const firstName string = "Cristiano"
	const lastName = "Ronaldo"
	fmt.Println(firstName + " " + lastName)

	// declare multi constants
	const (
		goals   = 9999
		assists = 8888
	)
	fmt.Println(goals)
	fmt.Println(assists)
}
