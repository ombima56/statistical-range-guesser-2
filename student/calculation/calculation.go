package calculation

import "math"

func LinearRegression(data []float64) (slope, intercept float64) {
	num := float64(len(data))
	var sumX, sumY, sumXY, sumX2 float64
	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}
	denominator := ((num * sumX2) - (sumX * sumX))
	if denominator == 0 {
		return
	}
	slope = (num*sumXY - sumX*sumY) / denominator
	intercept = (sumY - slope*sumX) / num
	return
}

func average(data []float64) float64 {
	n := len(data)

	if n == 0 {
		return 0
	}

	var sum float64
	for _, val := range data {
		sum += val
	}

	return sum / float64(n)
}

func variance(data []float64) (varianceVal float64) {
	n := len(data)

	av := average(data)

	var sqrDiffSUm float64
	for _, val := range data {
		diff := val - av
		sqrDiffSUm += diff * diff
	}

	varianceVal = sqrDiffSUm / float64(n)

	return
}

func StandardDeviation(data []float64) float64 {
	v := variance(data)
	return math.Sqrt(v)
}
