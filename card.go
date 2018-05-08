package deck // import "gophercise/deck"
import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//go:generate stringer -type=Suit,Rank

//Suit is one aspect of a card, it can be Spades, Hearts, Diamonds, Clubs, or Joker
type Suit uint8

//Rank is the numerical value of a card. it ranges from 1 - 13
type Rank uint8

const (
	spades Suit = iota
	hearts
	diamonds
	clubs
	joker
)

const (
	_ Rank = iota
	ace
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
)

//Card represents a playing card
type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

func (c Card) value() int {
	return int(c.Suit)*13 + int(c.Rank)
}

// New will return a new deck of cards as []Card
func New(opts ...func([]Card) []Card) []Card {
	deck := make([]Card, 52)
	for i := range deck {
		deck[i] = Card{Suit: Suit(i % 4), Rank: Rank((i % 13) + 1)}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck

}

//DefaultSort is a functional option for New()
//that will sort the deck by suit then rank
func DefaultSort(deck []Card) []Card {
	less := func(i, j int) bool {
		return deck[i].value() < deck[j].value()
	}
	sort.Slice(deck, less)
	return deck
}

//Sort is a functional option for New()
//that lets you sort the deck of card however you want
//by passing in a custom comparator of the format `func(i,j int) bool`
//see sort.Slice() docs for more info on the less func
func Sort(deck []Card, less func(int, int) bool) []Card {
	sort.Slice(deck, less)
	return deck
}

//Shuffle is a functional option for New()
//that will put your cards in a random order
//using the current time as the random seed
func Shuffle(deck []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled := make([]Card, len(deck))
	j := 0
	for _, i := range r.Perm(len(deck)) {
		shuffled[j] = deck[i]
		j++
	}
	return shuffled
}

//Jokers is a functional option for New() that will add an arbitrary number of jokers
//		Usage: New(Jokers(4)) will add 4 jokers to a new deck
func Jokers(n int) func([]Card) []Card {
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			deck = append(deck, Card{Suit: joker})
		}
		return deck
	}

}

//Filter is a functional option for New() that will remove cards from the deck
//the function that you pass into the function will return true for
//the cards you wish to remove
func Filter(f func(Card) bool) func([]Card) []Card {
	return func(deck []Card) []Card {
		var filteredDeck []Card
		for _, c := range deck {
			if !f(c) {
				filteredDeck = append(filteredDeck, c)
			}
		}
		return filteredDeck
	}
}

//Deck is a functional option for New()
//that will deplicate the current deck n number of times
//keep in mind that the order you put this in New will duplicate the deck with all of the previous options appllied
func Deck(n int) func([]Card) []Card {
	var ret []Card
	return func(deck []Card) []Card {
		for i := 0; i < n; i++ {
			ret = append(ret, deck...)
		}
		return ret
	}
}

//Merge Composes a single deck from any number of decks
func Merge(decks ...[]Card) []Card {
	var mergedDeck []Card
	for _, d := range decks {
		mergedDeck = append(mergedDeck, d...)
	}

	return mergedDeck
}
