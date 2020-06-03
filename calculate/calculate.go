package calculate

// CalcAvg calculates the average of a given slice
func CalcAvg(scores []int) float32 {
	var avg float32
	sum := sum(scores...)
	avg = float32(sum) / float32(len(scores))

	return avg
}

// CalcMedian clculates the median of a given slice
func CalcMedian(scores []int) int {
	var median int
	middle := len(scores) / 2
	if len(scores)%2 == 0 {
		median = (scores[middle-1] + scores[middle]) / 2
	} else {
		median = scores[middle]
	}
	return median
}

func sum(input ...int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}

	return sum
}
