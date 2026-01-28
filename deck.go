package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

// genDeck generates a deck of 52 cards with all combinations of suits and values.
// It returns a slice of strings where each string is a card in the format "value of suit".
func generateDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print prints out all the cards in the deck.
func (d *deck) print() {
	for _, card := range *d {
		fmt.Println(card)
	}
}

// deal takes a deck and a hand size and returns two slices of the deck.
// The first slice is the hand, and the second slice is the deck with the hand removed.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// shuffle shuffles the deck in place.
// It uses the current time as a source for randomness,
// and swaps each card with a random card in the deck.
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

// toString converts a deck into a string representation.
// It takes all the cards in the deck and joins them together with a comma separator.
// It returns a string representation of the deck.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saveToFile saves the deck to a file.
// It takes a filename as a string argument and saves the deck to the file.
// If there is an error saving the file, it returns the error.
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

// newDeckFromFile takes a filename as a string argument and returns a deck from the file.
// If there is an error reading the file, it logs the error and exits the program.
// It returns a deck created from the file contents, with each card separated by a comma.
// If the file is empty, it returns an empty deck.
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option #1 - Log the error and return a call to generateDeck()
		// Option #2 - Log the error and entirely quit the program

		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
