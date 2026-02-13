package engine

func CalculateWindow(currentIndex int, dailyNewWords int, totalWords int) (start int, end int) {
	if totalWords == 0 {
		return 0, -1
	}

	windowSize := dailyNewWords * 10

	end = currentIndex
	if end >= totalWords {
		end = totalWords - 1
	}

	start = end - windowSize + 1
	if start < 0 {
		start = 0
	}
	// Всегда начинать окно с первого слова
	if start > 0 {
		start = 0
	}

	return start, end
}