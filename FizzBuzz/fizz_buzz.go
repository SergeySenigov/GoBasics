package main

import (
	"fmt"
)

func main() {

	for message, i := "", 1; i <= 100; message, i = "", i+1 {

		if i%3 == 0 {
			message = message + "Fizz"
		}

		if i%5 == 0 {
			message = message + "Buzz"
		}

		if message == "" {
			message = fmt.Sprint(i)
		}

		fmt.Printf("%d: %s\n", i, message)
	}

}
