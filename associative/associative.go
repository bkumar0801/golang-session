package associative

import "fmt"

func ExampleMap()  {
	var x map[string]int
	x["key"] = 10
	fmt.Println(x)

	/*
	x := make(map[string]int)
	x["key"] = 10
	p := &x["key"]
	fmt.Println(x["key"])
	 */
}