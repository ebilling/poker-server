package api

import (
	"testing"

	"github.com/ebilling/poker-server/pb"
	"github.com/stretchr/testify/assert"
)

func TestCardNumbers(t *testing.T) {
	expectedSuit := []pb.Suit{
		pb.Suit_INVALID_SUIT,
		pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART, pb.Suit_HEART,
		pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE, pb.Suit_SPADE,
		pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND, pb.Suit_DIAMOND,
		pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB, pb.Suit_CLUB,
		pb.Suit_ANY, pb.Suit_ANY,
		pb.Suit_INVALID_SUIT,
	}
	expectedValue := []pb.CardValue{
		pb.CardValue_INVALID_CARD,
		pb.CardValue_ACE, pb.CardValue_TWO, pb.CardValue_THREE, pb.CardValue_FOUR, pb.CardValue_FIVE, pb.CardValue_SIX, pb.CardValue_SEVEN, pb.CardValue_EIGHT, pb.CardValue_NINE, pb.CardValue_TEN, pb.CardValue_JACK, pb.CardValue_QUEEN, pb.CardValue_KING,
		pb.CardValue_ACE, pb.CardValue_TWO, pb.CardValue_THREE, pb.CardValue_FOUR, pb.CardValue_FIVE, pb.CardValue_SIX, pb.CardValue_SEVEN, pb.CardValue_EIGHT, pb.CardValue_NINE, pb.CardValue_TEN, pb.CardValue_JACK, pb.CardValue_QUEEN, pb.CardValue_KING,
		pb.CardValue_ACE, pb.CardValue_TWO, pb.CardValue_THREE, pb.CardValue_FOUR, pb.CardValue_FIVE, pb.CardValue_SIX, pb.CardValue_SEVEN, pb.CardValue_EIGHT, pb.CardValue_NINE, pb.CardValue_TEN, pb.CardValue_JACK, pb.CardValue_QUEEN, pb.CardValue_KING,
		pb.CardValue_ACE, pb.CardValue_TWO, pb.CardValue_THREE, pb.CardValue_FOUR, pb.CardValue_FIVE, pb.CardValue_SIX, pb.CardValue_SEVEN, pb.CardValue_EIGHT, pb.CardValue_NINE, pb.CardValue_TEN, pb.CardValue_JACK, pb.CardValue_QUEEN, pb.CardValue_KING,
		pb.CardValue_JOKER, pb.CardValue_JOKER,
		pb.CardValue_INVALID_CARD,
	}
	for i := int32(0); i < 56; i++ {
		assert.Equal(t, expectedSuit[i], Suit(i))
		assert.Equal(t, expectedValue[i], Card(i))
	}
}

func TestCardValues(t *testing.T) {
}
