package engine

func CalculateWindow(currentIndex int, dailyNewWords int, totalWords int) (start int, end int) {
	if totalWords == 0 {
		return 0, -1
	}

	// Minimum words in session so we don't quit after just a few (e.g. 3).
	minWindowWords := dailyNewWords * 2
	if minWindowWords < 10 {
		minWindowWords = 10
	}

	end = currentIndex
	if end < minWindowWords-1 {
		end = minWindowWords - 1
	}
	if end >= totalWords {
		end = totalWords - 1
	}

	start = 0
	return start, end
}