# PACKAGE DOCUMENTATION
  This is a slightly modified output of `godoc github.com/Athulus/deck` 

```go
package deck
    import "github.com/Athulus/deck"
```

## FUNCTIONS

```go
func Deck(n int) func([]Card) []Card
```

Deck is a functional option for New() that will deplicate the current deck n number of times keep in mind that the order you put this in New will duplicate the deck with all of the previous options appllied

```go
func Filter(f func(Card) bool) func([]Card) []Card
```

Filter is a functional option for New() that will remove cards from the deck the function that you pass into the function will return true for the cards you wish to remove

```go
func Jokers(n int) func([]Card) []Card
```

Jokers is a functional option for New() that will add an arbitrary
number of jokers

Usage: New(Jokers(4)) will add 4 jokers to a new deck

## TYPES
``` go
type Card struct {
    Suit
    Rank
}
```
Card represents a playing card

```go
func DefaultSort(deck []Card) []Card
```

DefaultSort is a functional option for New() that will sort the deck by suit then rank

```go
func Merge(decks ...[]Card) []Card
```

Merge Composes a single deck from any number of decks

```go
func New(opts ...func([]Card) []Card) []Card
```
New will return a new deck of cards as []Card

```go
func Shuffle(deck []Card) []Card
```
Shuffle is a functional option for New() that will put your cards in a random order using the current time as the random seed

```go
func Sort(deck []Card, less func(int, int) bool) []Card
```

Sort is a functional option for New() that lets you sort the deck of card however you want by passing in a custom comparator of the format `func(i,j int) bool` see sort.Slice() docs for more info on the less func

```go
func (c Card) String() string
```

```go
type Rank uint8
```
Rank is the numerical value of a card. it ranges from 1 - 13

``` go
const (
    Ace Rank
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
```

```go
func (i Rank) String() string
```

```go
type Suit uint8
```
Suit is one aspect of a card, it can be Spades, Hearts, Diamonds, Clubs, or Joker

``` go
const (
    Spades Suit = iota
    Hearts
    Diamonds
    Clubs
    Joker
)
```

```go
func (i Suit) String() string
````
