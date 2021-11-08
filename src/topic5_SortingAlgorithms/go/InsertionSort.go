// QuickSort.go
// description: Implementation of in-place insertion sort algorithm
// details:
// A simple in-place insertion sort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Insertionsort)
// author [Oscar de Felice](https://github.com/oscar-defelice)

package main

import (
  "fmt"
)

// InsertionSort sorts the entire array
func InsertionSort(arr []int) (result []int) {
  for i := 1; i < len(arr); i++ {
		temporary := arr[i]
    j := i
		for ; j > 0 && arr[j-1] >= temporary; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temporary
	}
	return arr
}

func main() {
  slice := []int{5, 8, 1, 3, 7, 9, 2}
  fmt.Println("\n--- Unsorted --- \n\n", slice)
  fmt.Println("\n--- Sorted --- \n\n", InsertionSort(slice))
}
