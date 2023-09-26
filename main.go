package main

// To run multiple files of same package 'go run main.go state.go'
func main() {
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
