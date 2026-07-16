package deck

//go:generate stringer -type=Suit
type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade,Diamond,Club,Heart}

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
const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func New() []Card{
	var cards []Card
	for _, suit  := range suits {
    for rank := minRank; rank <= maxRank;rank++{
			cards = append(cards,Card{Suit: suit,Rank:rank})
		}  
	} 
	return cards
}

