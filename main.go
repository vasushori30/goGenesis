package main

// go mod init cards -> creates go.mod file

// To run multiple files of same package 'go run main.go state.go'
func main() {
	cards := newDeckFromFile("test")
	cards.shuffle()
	cards.print()
	// fmt.Println(cards.toString())
	// currentHandCards, remainingCards := deal(cards, 4)
	// currentHandCards.print()
	// remainingCards.print()
}
