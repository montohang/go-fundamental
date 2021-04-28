package main

// import "fmt"

// type Address struct {
// 	City, Province, Country string
// }

// // pointer di function
// func ChangeCountryToIndonesia(address *Address) {
// 	address.Country = "Indonesia"
// }

// // end pointer di function

// func main() {
// 	address1 := Address{
// 		"Bandung", "Jawa Barat", "Indonesia",
// 	}
// 	address2 := &address1
// 	address3 := &address1
// 	address2.City = "Padang"

// 	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

// 	alamat1 := new(Address)
// 	alamat2 := alamat1

// 	alamat2.Country = "Portugal"

// 	fmt.Println(*alamat1)
// 	fmt.Println(*alamat2)

// 	fmt.Println("===========================")

// 	fmt.Println(address1)
// 	fmt.Println(*address2)
// 	fmt.Println(*address3)

// 	// pointer di function
// 	var alamat = Address{
// 		City:     "Medan",
// 		Province: "Sumatera Utara",
// 		Country:  "",
// 	}
// 	fmt.Println("===========================")
// 	ChangeCountryToIndonesia(&alamat)
// 	fmt.Println(alamat)
// 	// end pointer di function
// }
