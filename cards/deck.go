package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), "|")
}

func (d deck) saveToFile(filename string) error {
	data := []byte(d.toString())
	const anyoneCanReadOrWrite = 0666
	return ioutil.WriteFile(filename, data, anyoneCanReadOrWrite)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(data), "|"))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	myRand := rand.New(source)
	for index := range d {
		newPosition := myRand.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
