package asciigenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomLine(t *testing.T) {
	assert.Equal(t, 1, 1)
}

/*
func TestGenerateDeckLength(t *testing.T) {
    deck := generateDeck()

    if len(deck) != 52 {
        t.Errorf("deck has wrong number of cards (should be 52): %d", len(deck))
        t.FailNow()
    }
}

func TestShuffleDeckLength(t *testing.T) {
    deck := shuffleDeck(generateDeck())

    if len(deck) != 52 {
        t.Errorf("deck has wrong number of cards (should be 52): %d", len(deck))
        t.FailNow()
    }
}

func TestShuffleDeckDuplicates(t *testing.T) {
    for i := 0; i <= 10; i++ {
        deck := shuffleDeck(generateDeck())
        var uniques []pokercard

        for _, v := range deck {
            for _, w := range uniques {
                if v == w {
                    t.Errorf("duplicate card found in shuffled deck: %+v %+v", v, w)
                    t.FailNow()
                }
            }
            uniques = append(uniques, v)
        }
        if len(uniques) != 52 {
            t.Errorf("not 52 %d", len(uniques))
            t.FailNow()
        }
    }
}

func TestDealLenth(t *testing.T) {
    deck := shuffleDeck(generateDeck())
    hand, deck := deal(deck)
    if len(hand) != 5 {
        t.Errorf("dealt hand has wrong number of cards (should be 5): %d", len(hand))
        for _, v := range hand {
            t.Errorf("dealt hand: %+v", v)
        }
        t.FailNow()
    }
}

func TestDealUnique(t *testing.T) {
    for i := 0; i < 10; i++ {
        deck := shuffleDeck(generateDeck())
        first, deck := deal(deck)
        second, deck := deal(deck)

        for _, v := range first {
            for _, w := range second {
                if v == w {
                    t.Errorf("two dealt hands have the same card -- first:\n %+v\n", first)
                    t.Errorf("two dealt hands have the same card -- second\n %+v\n", second)
                    t.Errorf("duplicates: %+v %+v", v, w)
                    t.FailNow()
                }
            }
        }
    }
}

func TestHighCard(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 3, true, },
        pokercard { 2, 4, true, },
        pokercard { 1, 5, true, },
        pokercard { 3, 7, true, },
    }

    high := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 3, true, },
        pokercard { 3, 4, true, },
        pokercard { 3, 9, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestOnePair(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 3, true, },
        pokercard { 3, 8, true, },
        pokercard { 3, 9, true, },
    }

    high := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 2, true, },
        pokercard { 3, 3, true, },
        pokercard { 3, 4, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestTwoPair(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 2, true, },
        pokercard { 3, 3, true, },
        pokercard { 3, 4, true, },
    }

    high := []pokercard {
        pokercard { 0, 3, true, },
        pokercard { 1, 4, true, },
        pokercard { 0, 4, true, },
        pokercard { 3, 5, true, },
        pokercard { 0, 5, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestThreeOfaKind(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 3, true, },
        pokercard { 1, 4, true, },
        pokercard { 0, 4, true, },
        pokercard { 3, 5, true, },
        pokercard { 0, 5, true, },
    }

    high := []pokercard {
        pokercard { 0, 1, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 2, true, },
        pokercard { 3, 2, true, },
        pokercard { 3, 1, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestStraight(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 9, true, },
        pokercard { 1, 2, true, },
        pokercard { 0, 1, true, },
        pokercard { 3, 2, true, },
        pokercard { 3, 8, true, },
    }

    high := []pokercard {
        pokercard { 0, 3, true, },
        pokercard { 1, 4, true, },
        pokercard { 0, 5, true, },
        pokercard { 2, 6, true, },
        pokercard { 3, 7, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestFlush(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 3, true, },
        pokercard { 1, 4, true, },
        pokercard { 0, 5, true, },
        pokercard { 2, 6, true, },
        pokercard { 3, 7, true, },
    }

    high := []pokercard {
        pokercard { 2, 1, true, },
        pokercard { 2, 3, true, },
        pokercard { 2, 8, true, },
        pokercard { 2, 9, true, },
        pokercard { 2, 10, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestFullHouse(t *testing.T) {
    low := []pokercard {
        pokercard { 2, 1, true, },
        pokercard { 2, 3, true, },
        pokercard { 2, 8, true, },
        pokercard { 2, 9, true, },
        pokercard { 2, 10, true, },
    }

    high := []pokercard {
        pokercard { 0, 2, true, },
        pokercard { 1, 2, true, },
        pokercard { 3, 2, true, },
        pokercard { 0, 4, true, },
        pokercard { 3, 4, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestFourOfaKind(t *testing.T) {
    low := []pokercard {
        pokercard { 0, 2, true, },
        pokercard { 1, 2, true, },
        pokercard { 3, 2, true, },
        pokercard { 0, 4, true, },
        pokercard { 3, 4, true, },
    }

    high := []pokercard {
        pokercard { 2, 3, true, },
        pokercard { 1, 3, true, },
        pokercard { 0, 3, true, },
        pokercard { 3, 3, true, },
        pokercard { 3, 8, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestStraightFlush(t *testing.T) {
    low := []pokercard {
        pokercard { 2, 3, true, },
        pokercard { 1, 3, true, },
        pokercard { 0, 3, true, },
        pokercard { 3, 3, true, },
        pokercard { 3, 8, true, },
    }

    high := []pokercard {
        pokercard { 2, 3, true, },
        pokercard { 2, 4, true, },
        pokercard { 2, 5, true, },
        pokercard { 2, 6, true, },
        pokercard { 2, 7, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}

func TestRoyalFlush(t *testing.T) {
    low := []pokercard {
        pokercard { 2, 3, true, },
        pokercard { 2, 4, true, },
        pokercard { 2, 5, true, },
        pokercard { 2, 6, true, },
        pokercard { 2, 7, true, },
    }

    high := []pokercard {
        pokercard { 3, 9, true, },
        pokercard { 3, 10, true, },
        pokercard { 3, 11, true, },
        pokercard { 3, 12, true, },
        pokercard { 3, 13, true, },
    }

    if scoreHand(low) > scoreHand(high) {
        t.Errorf("Low: %d, High: %d", scoreHand(low), scoreHand(high))
    }
}
*/
