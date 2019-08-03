package main

import "fmt"

func main() {

	const pi = 3

	add(1.00, 2.00)
	add(pi, 5.2)
	add(1, 2, 3)

	nums := []float64{1.00, 2.00, 8.12}

	add(nums...)
}

func add(nums ...float64) {

	fmt.Print("sum of", nums, "  ")

	total := 0.00

	for _, num := range nums {

		total = total + num
	}

	fmt.Println(total)

}
