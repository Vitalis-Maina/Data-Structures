package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Multiplication table: ", i)

		for k := 1; k <= 10; k++ {
			fmt.Printf(" %d * %d : %d\n", i, k, i*k)
		}
	}

}
