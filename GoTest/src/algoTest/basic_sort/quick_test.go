package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDivide(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 6, 7, 8, 2, 43, 1, 3, 2, 4, 9, 6, 85, 46}
	fmt.Println(divide(arr, 0, uint64(len(arr)-1)))
	fmt.Println(arr)
}
func divide(data []int, leftIndex, rightIndex uint64) (midL, midR uint64) {

	rand.Seed(time.Now().Unix())
	midIndex := leftIndex + rand.Uint64()%(rightIndex-leftIndex+1)
	mid := data[midIndex]
	p, pm, q := leftIndex-1, leftIndex, rightIndex+1
	for pm < q {
		if data[pm] < mid {
			// 交换
			tmp := data[pm]
			data[pm] = data[p+1]
			data[p+1] = tmp
			p++
			pm++
		} else if data[pm] == mid {
			pm++
		} else if data[pm] > mid {
			// 交换
			tmp := data[pm]
			data[pm] = data[q-1]
			data[q-1] = tmp
			q--
		}
	}
	return p + 1, pm - 1
}

func quickSort(data []int, leftIndex, rightIndex uint64) {
	if leftIndex >= rightIndex {
		return
	}
	l, r := divide(data, leftIndex, rightIndex)
	quickSort(data, leftIndex, l-1)
	quickSort(data, r+1, rightIndex)
}
func TestQuickSort(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6, 6, 7, 8, 2, 43, 1, 3, 2, 4, 9, 6, 85, 46}
	quickSort(arr, 0, uint64(len(arr)-1))
	fmt.Println(arr)
}
