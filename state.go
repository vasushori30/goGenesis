package main

import "fmt"

// Create a new type 'deck'
// which is slice of strings
type deck []string

// d = cards as in main.go we have cards.print()
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
