//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"math/rand"
	"sort"
)

type Suit int

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

type Rank int

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct{
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for suit := Spade; suit <= Heart; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	for _, option := range options {
		cards = option(cards)
	}
	return cards
}

func Sorted(cmp func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.SliceStable(cards, cmp(cards))
		return cards
	}
}

func DescendingOrder(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(card Card) int {
	return int(card.Suit) * int(King) + int(card.Rank)
}

func Shuffle(src rand.Source) func([]Card) []Card {
	return func(cards []Card) []Card {
		ret := make([]Card, len(cards))
		r := rand.New(src)
		for i, j := range r.Perm(len(cards)) {
			ret[i] = cards[j]
		}
		return ret
	}
}

func Jokers(count int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < count; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

func Filter(fn func(Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, card := range cards {
			if fn(card) {
				ret = append(ret, card)
			}
		}
		return ret
	}
}

func Decks(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var decks []Card
		for i := 0; i < n; i++ {
			decks = append(decks, cards...)
		}
		return decks
	}
}
