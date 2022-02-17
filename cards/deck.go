package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	result := deck{}

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			result = append(result, value+" of "+suit)
		}
	}
	return result
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
