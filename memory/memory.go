package memory

import "fmt"

var name string = "Terminator" //Stored in Data Segment

func AllocHeap() *int {
	p := new (int)  // Memory allocated in Heap but p will be in stack
	fmt.Println("AllocHeap ", p)
	return p
}

func AllocStack() *int {
	var a int // stored in stack
	fmt.Println("AllocStack ", &a)
	return &a;
}

func ReceiveOwnership()  {
	ptr := AllocHeap()
	fmt.Println("ReceiveOwnership (heap) ", ptr)
	fmt.Println("ReceiveOwnership (value) ", ptr)
	ptr = AllocStack()
	fmt.Println("ReceiveOwnership (stack) ", ptr)
}
