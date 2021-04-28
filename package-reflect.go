package main

// import (
// 	"fmt"
// 	"reflect"
// )

// // type Sample struct {
// // 	Name string
// // }

// // struct tag
// type Sample struct {
// 	Name string `required:"true" max:"10"`
// }

// // end struct tag

// func IsValid(data interface{}) bool {
// 	t := reflect.TypeOf(data)
// 	for i := 0; i < t.NumField(); i++ {
// 		field := t.Field(i)
// 		if field.Tag.Get("required") == "true" {
// 			value := reflect.ValueOf(data).Field(i).Interface()
// 			if value == "" {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func main() {
// 	// sample := Sample{"Harmon"}
// 	// sampleType := reflect.TypeOf(sample)
// 	// structField := sampleType.Field(0)

// 	// fmt.Println(sampleType.NumField())
// 	// fmt.Println(structField.Name)

// 	// struct tag
// 	sample := Sample{"Harmon"}
// 	sampleType := reflect.TypeOf(sample)
// 	structField := sampleType.Field(0)
// 	required := structField.Tag.Get("required")
// 	required1 := structField.Tag.Get("max")

// 	fmt.Println(required)
// 	fmt.Println(required1)

// 	// sample.Name = ""
// 	fmt.Println(IsValid(sample))
// 	// end struct tag
// }
