package main

import (
	"fmt"
	"math/rand/v2"
)

type Card struct {
	Rank string
	Suit string
}
type Deck []Card

type Player struct {
	Name string
	Hand Deck
	Ai   bool
}

//var deck []Card

func newd() Deck { // what a glorious function name
	cards := make(Deck, 0, 52)
	suits := []string{"S", "H", "D", "C"}
	ranks := []string{
		"A", "2", "3", "4", "5", "6", "7",
		"8", "9", "10", "J", "Q", "K",
	}
	for _, suit := range suits {
		for _, rank := range ranks {
			card := Card{Rank: rank, Suit: suit}
			cards = append(cards, card)
		}
	}
	//fmt.Printf("deck var: %v\n\nDeck length: %v\n", deck, len(deck))
	return cards
}

func shuf(cards Deck) {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
}

func deal(deck Deck, n int) (Deck, Deck, error) {
	if n < 0 || n > len(deck) {
		return nil, deck, fmt.Errorf("Can't deal %d cards from the deck of size %d\n", n, len(deck))
	}
	hand := make(Deck, n)
	copy(hand, deck[:n])
	rem := deck[n:]
	return hand, rem, nil
}
