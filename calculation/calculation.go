package calculation


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