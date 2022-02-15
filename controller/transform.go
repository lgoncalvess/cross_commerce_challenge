package controller

//Algoritmo de ordenação Merge Sort
func MergeSort(items []float64) []float64 {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []float64, b []float64) []float64 {
	final := []float64{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
