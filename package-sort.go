package main

// import (
// 	"fmt"
// 	"sort"
// )

// type User struct {
// 	Name string
// 	Age  int
// }

// type UserSlice []User

// func (userSlice UserSlice) Len() int {
// 	return len(userSlice)
// }

// func (userSlice UserSlice) Less(i, j int) bool {
// 	return userSlice[1].Age < userSlice[j].Age
// }

// func (userSlice UserSlice) Swap(i, j int) {
// 	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
// }

// func main() {
// 	users := []User{
// 		{"Harmon", 30},
// 		{"Mon", 35},
// 		{"Tohang", 31},
// 		{"montohang", 25},
// 	}

// 	sort.Sort(UserSlice(users))

// 	fmt.Println(users)
// }
