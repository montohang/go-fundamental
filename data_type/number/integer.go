package main

import "fmt"

func DataTypeIntegers() {
	var (
		// integers
		dataTypeInt   int
		dataTypeInt8  int8  = 127
		dataTypeInt16 int16 = 32767
		dataTypeInt32 int32 = 2147483647
		dataTypeInt64 int64 = 9223372036854775807
		// unsigned integers
		dataTypeUint   uint
		dataTypeUint8  uint8  = 255
		dataTypeUint16 uint16 = 65535
		dataTypeUint32 uint32 = 4294967295
		dataTypeUint64 uint64 = 18446744073709551615
	)

	fmt.Println(dataTypeInt, dataTypeInt8, dataTypeInt16, dataTypeInt32, dataTypeInt64)
	fmt.Println(dataTypeUint, dataTypeUint8, dataTypeUint16, dataTypeUint32, dataTypeUint64)
}
