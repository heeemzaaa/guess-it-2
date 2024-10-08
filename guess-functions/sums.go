package guess

// this function calculates all the sums we need in the two functions below
func Sums(data []float64) (float64, float64, float64, float64, float64) {
	var sumx float64
	var sumy float64
	var sumxy float64
	var sumx2 float64
	var sumy2 float64
	for i := 0; i < len(data); i++ {
		sumx += float64(i)
	}
	for i := 0; i < len(data); i++ {
		sumy += data[i]
	}
	for i := 0; i < len(data); i++ {
		sumxy += data[i] * float64(i)
	}
	for i := 0; i < len(data); i++ {
		sumx2 += float64(i) * float64(i)
	}
	for i := 0; i < len(data); i++ {
		sumy2 += data[i] * data[i]
	}

	return sumx, sumy, sumxy, sumx2, sumy2
}
