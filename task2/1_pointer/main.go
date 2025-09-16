package main

import "fmt"

// task1.1
func intPointer(i *int) {
	*i += 10
}

// task1.2
func sliceIntPointer(i *[]int) {
	for k, v := range *i {
		(*i)[k] = v * 2
	}
}

func main() {
	// task1.1
	// a := 10
	// intPointer(&a)
	// fmt.Println("intPointer", a)

	// task1.2
	b := []int{1, 2, 3}
	sliceIntPointer(&b)
	fmt.Println("sliceIntPointer", b)

}
