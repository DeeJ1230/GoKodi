package main

import "fmt"

func print(msg string) {
	fmt.Println(msg)
}

func main() {

	words := [...]string{"lets", "do", "this", "now"}

	for _, n := range words {
		fmt.Printf("Word: %s\n", n)
	}

	fmt.Printf("Word #2: %s\n", words[1])

	for n := 1; n < 10; n++ {
		print("Hello")
	}
}
