package main

func isHeap(heap []int) bool {
	for i := 0; i < len(heap); i++ {
		left := 2*i + 1
		right := left + 1
		if left < len(heap) && heap[left] < heap[i] {
			return false
		}
		if right < len(heap) && heap[right] < heap[i] {
			return false
		}
	}
	return true
}
