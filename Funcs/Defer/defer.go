package main

import "fmt"

var Global = 5

func use() {

	defer func(val int) {
		Global = val
	}(Global)

	Global = 4
	fmt.Println("Changed", Global)
}

func main() {

	fmt.Println("Start", Global)
	use()
	fmt.Println("Final", Global)
}
