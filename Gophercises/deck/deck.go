package deck

import "fmt"

//go:generate stringer -type=Suit

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)
//go:generate stringer -type=Rank
type Rank uint8

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


type Card struct {
	//suit
	//rank
	Suit 
	Rank

}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}



func New() []Card {
	var cards []Card
	//for each suit
	//for each rank
	// add card{suit,rank} to cards
}