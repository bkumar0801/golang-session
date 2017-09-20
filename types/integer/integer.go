/*
(1) Go's integer types are: uint8, uint16, uint32, uint64,
                           int8, int16, int32 and int64.
(2) In addition there two alias types: byte which is the same
    as uint8 and rune which is the same as int32.
(3) There are also 3 machine dependent integer types: uint, int and uintptr.
*/

package integer

import (
	"fmt"
	"unsafe"
)

func WhichIsBetter()   {
	a := 20
	var b uint8 = 20
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("a ", unsafe.Sizeof(a))
	fmt.Println("b ", unsafe.Sizeof(b))
}

func WhatIsExpected(value int) uint8 {
	var a uint8 = uint8(value)
	fmt.Println("a ", a)
	return a
}

func UpCast(value uint8)  int {
	var a int = int(value)
	fmt.Println("a ", a)
	return a;
}
