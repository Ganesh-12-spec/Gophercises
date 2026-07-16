package main

import (
	"deckofcards/deck"
	"fmt"
)

func main() {
   cards := deck.New()
	 fmt.Println(len(cards))
}