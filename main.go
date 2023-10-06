package main

import "fmt"

// To run multiple files of same package 'go run main.go state.go'
func main() {
	cards := newDeckFromFile("test1")
	fmt.Println(cards)
	// fmt.Println(cards.toString())
	// currentHandCards, remainingCards := deal(cards, 4)
	// currentHandCards.print()
	// remainingCards.print()
}
