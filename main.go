package main

// To run multiple files of same package 'go run main.go state.go'
func main() {
	//Slice can have any number of elements
	//Slice element can only have 1 type
	cards := deck{"Ace of Diamonds", newCard()}
	//append in the background create a new slice and do not modify the existing slice when we append new element
	cards = append(cards, "Six of Hearts")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
