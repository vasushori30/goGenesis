package main

import "fmt"

func main() {
	//Slice can have any number of elements
	//Slice element can be of only 1 type
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Hearts")
	//append in the background create a new slice and do not modify the existing slice when we append new element
	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
