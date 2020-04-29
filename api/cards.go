package api

import (
	"github.com/ebilling/poker-server/pb"
)

// Suit returns the suit for an integer representation of a card
func Suit(card int32) pb.Suit {
	if card < 1 || card > 54 {
		return pb.Suit_INVALID_SUIT
	}
	if card > 52 {
		return pb.Suit_ANY
	}
	return pb.Suit(1 + (card-1)/13)
}

// Card returns the CardValue for an integer representation of a card
func Card(card int32) pb.CardValue {
	if card < 1 || card > 54 {
		return pb.CardValue_INVALID_CARD
	}
	if card > 52 {
		return pb.CardValue_JOKER
	}
	return pb.CardValue(1 + (card-1)%13)
}
