package guess

import "math"

// this function calculates the Pearson correlation coefficient
func Pearson_Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2 float64, data []float64) float64 {
	length := float64(len(data))
	var r float64
	holder1 := (length * sumxy) - (sumx * sumy)
	holder2 := ((length * sumx2) - (sumx * sumx)) * ((length * sumy2) - (sumy * sumy))
	r = holder1 / math.Sqrt(holder2)
	return r
}
