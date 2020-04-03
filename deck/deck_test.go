package deck_test

import (
	"fmt"
	"github.com/ambye85/gophercises/deck"
	"math/rand"
	"sort"
	"testing"
)

const seed = 0

var suits = []deck.Suit{deck.Spade, deck.Diamond, deck.Club, deck.Heart}
var ranks = []deck.Rank{deck.Ace, deck.Two, deck.Three, deck.Four, deck.Five, deck.Six, deck.Seven, deck.Eight, deck.Nine, deck.Ten, deck.Jack, deck.Queen, deck.King}

func ExampleCard() {
	fmt.Println(deck.Card{Rank: deck.Ace, Suit: deck.Heart})
	fmt.Println(deck.Card{Rank: deck.Two, Suit: deck.Spade})
	fmt.Println(deck.Card{Rank: deck.Ten, Suit: deck.Diamond})
	fmt.Println(deck.Card{Rank: deck.King, Suit: deck.Club})
	fmt.Println(deck.Card{Suit: deck.Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Ten of Diamonds
	// King of Clubs
	// Joker
}

func standardDeck() []deck.Card {
	deckSize := 52
	expected := make([]deck.Card, deckSize, deckSize)
	for i, suit := range suits {
		for j, rank := range ranks {
			expected[(i*len(ranks))+j] = deck.Card{Suit: suit, Rank: rank}
		}
	}
	return expected
}

func decksEqual(actual []deck.Card, expected []deck.Card, t *testing.T) {
	if len(actual) != len(expected) {
		t.Fatalf("got %d, wanted %d", len(actual), len(expected))
	}

	for i, got := range actual {
		if want := expected[i]; want != got {
			t.Fatalf("got %+v, wanted %+v", got, want)
		}
	}
}

func TestCreateNewDeckWithDefaultSortOrder(t *testing.T) {
	cards := deck.New()

	expected := standardDeck()

	decksEqual(cards, expected, t)
}

func TestCreateDeckWithCustomSorting(t *testing.T) {
	cards := deck.New(deck.Sorted(deck.DescendingOrder))

	expected := standardDeck()
	sort.SliceStable(expected, func(i, j int) bool {
		return (int(expected[i].Suit)*len(ranks))+int(expected[i].Rank) > (int(expected[j].Suit)*len(ranks))+int(expected[j].Rank)
	})

	decksEqual(cards, expected, t)
}

func TestShufflesCards(t *testing.T) {
	src := rand.NewSource(seed)
	cards := deck.New(deck.Shuffle(src))

	expected := deck.Card{Rank: deck.Two, Suit: deck.Heart}
	if cards[0] != expected {
		t.Fatalf("got %+v, wanted %+v", cards[0], expected)
	}
}

func TestJokers(t *testing.T) {
	cards := deck.New(deck.Jokers(2))

	if len(cards) != 54 {
		t.Fatalf("got %d, wanted %d", len(cards), 54)
	}

	if joker := cards[52]; joker.Suit != deck.Joker {
		t.Fatalf("got %s, wanted %s", joker, deck.Joker)
	}

	if joker := cards[53]; joker.Suit != deck.Joker {
		t.Fatalf("got %s, wanted %s", joker, deck.Joker)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card deck.Card) bool {
		return card.Suit == deck.Heart || card.Suit == deck.Diamond
	}

	cards := deck.New(deck.Filter(filter))
	for _, card := range cards {
		if card.Suit == deck.Spade || card.Suit == deck.Club {
			t.Fatalf("should not have got a %s", card.Suit)
		}
	}
}

func TestMultipleDecks(t *testing.T) {
	numDecks := 2
	cards := deck.New(deck.Decks(numDecks))

	counts := make(map[deck.Card]int)
	for _, card := range cards {
		counts[card]++
	}

	for card, count := range counts {
		if count != numDecks {
			t.Errorf("%s: got %d, wanted %d", card, count, numDecks)
		}
	}
}

func TestMultipleDecksWithJokersOptionBeforeDecksOption(t *testing.T) {
	numDecks := 2
	numJokers := 2
	cards := deck.New(deck.Jokers(numJokers), deck.Decks(numDecks))

	count := 0
	for _, card := range cards {
		if card.Suit == deck.Joker {
			count++
		}
	}

	if count != numDecks*numJokers {
		t.Fatalf("got %d, wanted %d", count, numDecks*numJokers)
	}
}

func TestMultipleDecksWithJokersOptionAfterDecksOption(t *testing.T) {
	numDecks := 2
	numJokers := 2
	cards := deck.New(deck.Decks(numDecks), deck.Jokers(numJokers))

	count := 0
	for _, card := range cards {
		if card.Suit == deck.Joker {
			count++
		}
	}

	if count != numJokers {
		t.Fatalf("got %d, wanted %d", count, numJokers)
	}
}
