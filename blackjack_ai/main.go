package main

import (
	"fmt"
	"github.com/ambye85/gophercises/blackjack/blackjack"
)

func main() {
	game := blackjack.New(blackjack.Options{
		Decks:           3,
		Hands:           10,
		BlackjackPayout: 1.5,
	})
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
