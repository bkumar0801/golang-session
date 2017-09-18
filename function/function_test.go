package function

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	fmt.Println(Add(1,2,3))
}

func TestMakeEvenGenerator(t *testing.T) {
	nextEven := MakeEvenGenerator()
	fmt.Println(nextEven)
}
