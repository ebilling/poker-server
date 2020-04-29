package api

import (
	"github.com/chehsunliu/poker"
	"github.com/ebilling/poker-server/pb"
)

// Hand holds the information about a single game at a single table.
type Hand struct {
	Deck      *poker.Deck
	Blind     int32
	PlayQueue []*pb.Play
}

// NewHand constructs a poker Hand with a brand new deck of cards
func NewHand(bigBlind int32, players []*pb.Player) *Hand {
	return &Hand{
		Deck:  poker.NewDeck(),
		Blind: bigBlind,
	}
}

func (h *Hand) readyToAdvance(table *pb.TableStatus) bool {
	for _, player := range table.GetPlayers() {
		table.
	}
	return true
}

func Advance(table *pb.TableStatus) {

}
