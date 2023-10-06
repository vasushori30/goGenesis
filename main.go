package main

// To run multiple files of same package 'go run main.go state.go'
func main() {
	cards := newDeck()
	cards.saveToFile("test")
	// fmt.Println(cards.toString())
	// currentHandCards, remainingCards := deal(cards, 4)
	// currentHandCards.print()
	// remainingCards.print()
}
