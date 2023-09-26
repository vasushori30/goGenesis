package main

import "fmt"

func main() {
	//Slice can have any number of elements
	//Slice element can only have 1 type
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Hearts")
	//append in the background create a new slice and do not modify the existing slice when we append new element
	for index, card := range cards {
		fmt.Println(index, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
