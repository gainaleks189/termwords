package engine

// CalculateWindow returns the sliding review window (start, end).
// Each word stays in the window for 10 days; window size = dailyNewWords * 10.
// Old words drop out automatically when start increases.
func CalculateWindow(currentIndex int, dailyNewWords int, totalWords int) (start int, end int) {
	if totalWords == 0 {
		return 0, -1
	}

	windowSize := dailyNewWords * 10
	end = currentIndex
	if end >= totalWords {
		end = totalWords - 1
	}
	start = currentIndex - windowSize + 1
	if start < 0 {
		start = 0
	}
	return start, end
}