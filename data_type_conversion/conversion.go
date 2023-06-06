package main

import "fmt"

func main() {
	var dataTypeInt32 int32 = 32768
	var dataTypeInt64 = int64(dataTypeInt32)
	fmt.Println(dataTypeInt32)
	fmt.Println(dataTypeInt64)

	var name = "Cristiano Ronaldo"
	var getOneCharacter = name[0]
	var eString = string(getOneCharacter)
	fmt.Println(name)
	fmt.Println(getOneCharacter)
	fmt.Println(eString)

}
