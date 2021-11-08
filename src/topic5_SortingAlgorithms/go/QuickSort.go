// QuickSort.go
// description: Implementation of in-place quicksort algorithm
// details:
// A simple in-place quicksort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Quicksort)
// author [Oscar de Felice](https://github.com/oscar-defelice)

package main

import (
  "fmt"
)

// QuickSort Sorts the entire array recursively
func QuickSort(arr []int) (result []int) {
  if len(arr) <= 1 {
    return arr
  }

  left, right := make([]int, 0), make([]int, 0)
  pivot := arr[0]

  for i, _ := range arr {
      if arr[i] < pivot {
        left = append(left, arr[i])
      } else if arr[i] > pivot {
        right = append(right, arr[i])
      }
  fmt.Println(i, "left: ", left, "\n right: ", right)
  }
  sort_left, sort_right := QuickSort(left), QuickSort(right)
  result = append(sort_left, pivot)
  result = append(result, sort_right...)
  return result
}

func main()  {
  fmt.Println("\n--- Unsorted --- \n\n", slice)
  fmt.Println("\n--- Sorted --- \n\n", QuickSort(slice))
}
