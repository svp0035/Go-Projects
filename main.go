package main

import "github.com/svp0035/calculator/operations"

func main() {
	operations.AddInt8(120, 110)      // the outputted number will be wrong because it is out of range of int8
	operations.AddInt16(30000, 30000) // the outputted number will be wrong because it is out of range of int8
	operations.AddInt32(120, 110)
	operations.AddInt64(120, 110)
	operations.AddUInt8(120, 110)
	operations.AddUInt16(120, 110)
	operations.AddUInt32(120, 110)
	operations.AddUInt64(120, 110)
	operations.AddFloat32(120.0230, 110.3218)
	operations.AddFloat64(120.3443, 110.2340)

	operations.SubInt8(120, 110)
	operations.SubInt16(120, 110)
	operations.SubInt32(120, 110)
	operations.SubInt64(120, 110)
	operations.SubUInt8(120, 110)
	operations.SubUInt16(120, 110)
	operations.SubUInt32(120, 110)
	operations.SubUInt64(120, 110)
	operations.SubFloat32(120, 110)
	operations.SubFloat64(120, 110)

	operations.MultiplyInt8(120, 110)
	operations.MultiplyInt16(120, 110)
	operations.MultiplyInt32(120, 110)
	operations.MultiplyInt64(120, 110)
	operations.MultiplyUInt8(120, 110)
	operations.MultiplyUInt16(120, 110)
	operations.MultiplyUInt32(120, 110)
	operations.MultiplyUInt64(120, 110)
	operations.MultiplyFloat32(120, 110)
	operations.MultiplyFloat64(120, 110)

	operations.DivideInt8(120, 110)
	operations.DivideInt16(120, 110)
	operations.DivideInt32(120, 110)
	operations.DivideInt64(120, 110)
	operations.DivideUInt8(120, 110)
	operations.DivideUInt16(120, 110)
	operations.DivideUInt32(120, 110)
	operations.DivideUInt64(120, 110)
	operations.DivideFloat32(120, 110)
	operations.DivideFloat64(120, 110)
}
