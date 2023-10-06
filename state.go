package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type 'deck'
// which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// d = cards as in main.go we have cards.print()
// deck as a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// paramaters as a arguments
func deal(d deck, handSize int) (deck, deck) {
	//slice range syntax d[startIndexIncluding:upToNotIncluding]
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// This filename will create a new file and write the deck to it with 0666 permission
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read from file
func newDeckFromFile(filename string) deck {
	// bs = byte slice
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// This is how we convert byte slice to string
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		// Generate a random number
		newPosition := rand.Intn(len(d) - 1)
		// Swap the position
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
