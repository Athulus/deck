package deck

import (
	"reflect"
	"testing"
)

func TestCard_String(t *testing.T) {
	type fields struct {
		Suit Suit
		Rank Rank
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				Suit: tt.fields.Suit,
				Rank: tt.fields.Rank,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want []Card
	}{
		{name: "test", want: []Card{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultSort(t *testing.T) {
	tests := []struct {
		name string
		want []Card
	}{
		{name: "test", want: []Card{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(DefaultSort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		want []Card
	}{
		{name: "test", want: []Card{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(DefaultSort, Shuffle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJokers(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []Card
	}{
		{name: "test", args: 4, want: []Card{}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(Jokers(tt.args))
			numberOfJokers := 0
			for _, c := range got {
				if c.Suit == joker {
					numberOfJokers++
				}
			}
			if numberOfJokers != tt.args {
				t.Errorf("wanted %v jokers, got %v", tt.args, numberOfJokers)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	filter := func(c Card) bool {
		if c.Suit != hearts {
			return false
		}
		return true
	}
	testDeck := New(Filter(filter))
	for _, c := range testDeck {
		if c.Suit == hearts {
			t.Error("found a heart where there should be none")
		}
	}
}
