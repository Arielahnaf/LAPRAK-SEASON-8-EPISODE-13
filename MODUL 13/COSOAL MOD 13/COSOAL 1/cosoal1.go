package main

import "fmt"

func main() {
	var word string
	var repititions int
	fmt.Scan(&word, &repititions)
	for i := 0; i < repititions; i++ {
		fmt.Println(word)
	}
}
