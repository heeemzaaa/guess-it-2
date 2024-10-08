package guess

// this function predicts the range of the next number
func Guess_it(data []float64) (float64, float64) {
	start := len(data) - 4
	if start < 0 {
		start = 0
	}

	slice := data[start:]
	sumx, sumy, sumxy, sumx2, sumy2 := Sums(slice)
	min, max := Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2, slice)
	r := Pearson_Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2, slice) * 25

	min -= r
	max += r

	if min < 0 {
		min *= -1
	}
	if max < 0 {
		max *= -1
	}
	if min > max {
		min, max = max, min
	}
	if min+r < max {
		min = max + r
		if min > max {
			min = max - r
		}
	} else {
		min = max - r
	}
	return min, max
}
