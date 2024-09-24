package main

func getMode(arr []float64) (mode int) {
	//	Create a map and populated it with each value in the slice
	//	mapped to the number of times it occurs
	countMap := make(map[float64]float64)
	for _, value := range arr {
		countMap[value] += 1
	}

	//	Find the smallest item with greatest number of occurance in
	//	the input slice
	max := float64(0)
	for _, key := range arr {
		freq := countMap[key]
		if freq > max {
			mode = int(key)
			max = freq
		}
	}
	return
}

func getAvg(arr []float64) (avg int) {
	total := float64(0)
	for _, number := range arr {
		total = total + number
	}
	return int(total / float64(len(arr)))
}
