// Package sorter provides a function to create a sorted index of data elements.
// It is useful for iteration over slices of arbitrary data structures in sorted order.
// Instead of sorting actual data it produces a sorted slice of data indices.
package sorter

// CreateIndex produces an slice of data element indices sorted using given function.
// The dataLen parameter is the length of data to sort.
// The lessFunc parameter is a function that is used to compare two elements with indexes i and j.
func CreateIndex(dataLen int, lessFunc func(i, j int) bool) []int {
	index := make([]int, dataLen)
	for i := 0; i < dataLen; i++ {
		index[i] = i
	}
	quicksort(index, 0, dataLen-1, lessFunc)
	return index
}

func quicksort(index []int, start, stop int, lessFunc func(i, j int) bool) {
	if start >= stop {
		return
	}
	i := start
	j := stop
	k := index[start+(stop-start)/2]
	for i <= j {
		for lessFunc(index[i], k) {
			i++
		}
		for lessFunc(k, index[j]) {
			j--
		}
		if i < j {
			index[i], index[j] = index[j], index[i]
		}
		if i <= j {
			i++
			j--
		}
	}
	if i < stop {
		quicksort(index, i, stop, lessFunc)
	}
	if j > start {
		quicksort(index, start, j, lessFunc)
	}
}
