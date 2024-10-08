package guess

// this function calculates the linear regression line
func Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2 float64, data []float64) (float64, float64) {
	var a float64
	var b float64
	length := len(data)

	b = ((float64(length) * sumxy) - (sumx * sumy)) / ((float64(length) * sumx2) - (sumx * sumx))
	a = (sumy / float64(length)) - ((b * sumx) / float64(length))
	return b, a
}
