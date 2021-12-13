package asciigenerator

func RandomLine() (string, error) {
	return "TODO", nil
}

/*
// TODO: repurpose

type pokercard struct {
    suit int
    value int
    dealt bool
}

type byValue []pokercard

func (a byValue) Len() int {
    return len(a)
}

func (a byValue) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a byValue) Less(i, j int) bool {
    return a[i].value < a[j].value
}

func shuffleDeck(unshuffled []pokercard) []pokercard {
	deck := make([]pokercard, len(unshuffled))
	order := rand.Perm(len(unshuffled))
	for i, v := range order {
		deck[i] = unshuffled[v]
	}

    return deck
}

func generateDeck() []pokercard {
    var deck []pokercard
    for i := 0; i < 13; i++ {
        for j := 0; j <= 3; j++ {
            deck = append(deck, pokercard {
                j,
                i,
                false,
            })
        }
    }

    return deck
}

func getUndealtCard(deck []pokercard) pokercard {
    rand.Seed(time.Now().Unix())
	for _, i := range rand.Perm(len(deck)) {
		if !deck[i].dealt {
			deck[i].dealt = true
			return deck[i]
		}
	}
    return getUndealtCard(deck)
}

func deal(deck []pokercard) ([]pokercard, []pokercard) {
    hand := make([]pokercard, 5)
    for i := 0; i <= 4; i++ {
        hand[i] = getUndealtCard(deck)
    }
    return hand, deck
}

func translateSuit(suit int) string {
	translation := []string {
		"hearts",
		"diamonds",
		"clubs",
		"spades",
	}
	return " of " + translation[suit]
}

func translateHand(hand []pokercard) string {
    var r string
    for _, v := range hand {
        if v.value == 12 {
            r += "king"
        } else if v.value == 11 {
            r += "queen"
        } else if v.value == 10 {
            r += "jack"
        } else if v.value == 0 {
            r += "ace"
        } else {
            r += strconv.Itoa(v.value+1)
        }
        r+= translateSuit(v.suit) +
        ", "
    }
    return r
}

func royal(hand []pokercard) int {
    for i := len(hand)-1; i >= 0; i-- {
        if i+9 != hand[i].value {
            return 0
        }
    }
    return 1000
}

func flush(hand []pokercard) int {
    suit := -1
    for _, v := range hand {
        if suit == -1 {
            suit = v.suit
            continue
        }
        if  suit != v.suit {
            return 0
        }
    }

    return 600
}

func straight(hand []pokercard) int {
    value := -1
    for _, v := range hand {
        if value == -1 {
            value = v.value
            continue
        } else if  v.value != value + 1 {
            return 0
        }
        value = v.value
    }
    return 500
}

func pairs(hand []pokercard) int {
	var matches = make(map[int]int)

	for _, v := range hand {
		matches[v.value]++
	}

    threeKind := false
    fourKind := false
    twoPair :=  false
    onePair := false

	for _, v := range matches {
        if v == 2 {
	        if onePair == true {
	            twoPair = true
            } else {
                onePair = true
            }
        } else if v == 3 {
            threeKind = true
        } else if v == 4 {
            fourKind = true
        }
    }

    if onePair && threeKind {
        return 700
    } else if fourKind {
        return 800
    } else if threeKind {
        return 200
    } else if twoPair {
        return 100
    } else if onePair {
        return 50
    } else {
	    return 0
    }
}

func highest(hand []pokercard) int {
    return hand[len(hand)-1].value
}

func scoreHand(hand []pokercard) int {
    return royal(hand) +
        flush(hand) +
        straight(hand) +
        pairs(hand) +
        highest(hand)
}

func declareWinner(player []pokercard, bot []pokercard) string {
    if scoreHand(player) >= scoreHand(bot) {
        return "and you win!"
    } else {
        return "and I win!"
    }
}

func compareHands(first []pokercard, second []pokercard) string {
    return "Your hand is: " +
            translateHand(first) +
            "my hand is: " +
            translateHand(second) +
            declareWinner(first, second)
}

func playPoker() string {
    deck := shuffleDeck(generateDeck())
    player, deck := deal(deck)
    bot, deck := deal(deck)
    sort.Sort(byValue(player))
    sort.Sort(byValue(bot))
    return compareHands(bot, player)
}
*/
