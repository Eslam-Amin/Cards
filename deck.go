package main

import "fmt"

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

// receiver function to print the deck
func (d *deck) print() {
	for _, card := range *d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
