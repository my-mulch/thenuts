package poker

func IsStraight(ranks RankCounter) bool {
	if len(ranks) != 5 {
		return false
	}

	numCards := 0

	_, hasLowAce := ranks[Ace]
	if hasLowAce {
		numCards = 1
	}

	for rank := Two; rank <= Ace; rank++ {
		_, hasRank := ranks[rank]

		// We have a straight
		if hasRank && (numCards+1) == 5 {
			return true
		}

		// Keep going, we haven't missed yet
		if hasRank {
			numCards += 1
			continue
		}

		// We missed, but we don't have any cards yet
		if numCards == 0 {
			continue
		}

		// We missed, but we have a low ace. Maybe the straight is royal
		if hasLowAce && rank == Two {
			numCards = 0
			continue
		}

		return false
	}

	return false
}

func IsNOfAKind(ranks RankCounter, n int) bool {
	for _, count := range ranks {
		if count == n {
			return true
		}
	}

	return false
}

func IsTwoPair(ranks RankCounter) bool {
	numPairs := 0

	for _, count := range ranks {
		if count == 2 {
			numPairs += 1
		}
	}

	return numPairs == 2
}