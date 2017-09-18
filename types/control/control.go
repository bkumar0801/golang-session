package control

import "fmt"

func ForLoop() {
	i := 1  // catch problem
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
}

func IfBlock()  {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func SwitchBlock(i int)  {
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}
}


