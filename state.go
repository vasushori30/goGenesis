package main

import "fmt"

// Create a new type 'deck'
// which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"One", "Diamonds", "Hearts", "Club"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// d = cards as in main.go we have cards.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
