package main

import "fmt"

func main() {
	fmt.Println("Enter Input")

	var input int

	fmt.Scanln(&input)
	printTable(input)
}

func printTable(input int) {
	var i int
	j := 1
	count := (input * 10) + 1
	for i = input; i < count; i += input {
		fmt.Println(input, "*", j, "=", i)
		j += 1
	}
}
