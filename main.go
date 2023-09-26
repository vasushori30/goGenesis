package main

// To run multiple files of same package 'go run main.go state.go'
func main() {
	cards := newDeck()
	currentHandCards, remainingCards := deal(cards, 4)
	currentHandCards.print()
	remainingCards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
