package main

func colorMap(hand []Card) map[int]int {
	returnMap := make(map[int]int)
	for _, card := range hand {
		val, ok := returnMap[card.family]
		if ok {
			returnMap[card.family] = val + 1
		} else {
			returnMap[card.family] = 1
		}
	}
	return returnMap
}

func valueMap(hand []Card) map[int]int {
	returnMap := make(map[int]int)
	for _, card := range hand {
		val, ok := returnMap[card.value]
		if ok {
			returnMap[card.value] = val + 1
		} else {
			returnMap[card.value] = 1
		}
	}
	return returnMap
}

func CompareHands(handA []Card, handB []Card) int {

	aValue := valueMap(handA)
	bValue := valueMap(handB)

	aColor := colorMap(handA)
	bColor := colorMap(handB)

	// couleur
	for c := 0; c < 4; c++ {
		aCount, aOk := aColor[c]
		bCount, bOk := bColor[c]
		aPair := aOk && aCount == 5
		bPair := bOk && bCount == 5
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// suite

	// full

	// brelan
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 3
		bPair := bOk && bCount == 3
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// double paire
	aPairs := make([]int, 0)
	bPairs := make([]int, 0)
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 2
		bPair := bOk && bCount == 2
		if aPair {
			aPairs = append(aPairs, val)
		}
		if bPair {
			bPairs = append(bPairs, val)
		}
	}
	if len(aPair) == 2 && len(bPairs) < 2 {
		return 1
	}
	if len(bPair) == 2 && len(aPairs) < 2 {
		return -1
	}

	// paire
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 2
		bPair := bOk && bCount == 2
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}

	// carte la plus haute
	for val := 12; val >= 0; val-- {
		aCount, aOk := aValue[val]
		bCount, bOk := bValue[val]
		aPair := aOk && aCount == 1
		bPair := bOk && bCount == 1
		if aPair && !bPair {
			return 1
		}
		if bPair && !aPair {
			return -1
		}
	}
	return 0
}
