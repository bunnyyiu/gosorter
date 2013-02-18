package mergesortp

import (
  "runtime"
  "fmt"
)

func merge(left, right []int) (results []int){
  lenL := len(left)
  lenR := len(right)
  size := lenL + lenR
  results = make([]int, 0, size)
  for lenL > 0 && lenR > 0 {
    if left[0] <= right[0] {
      results = append(results, left[0])
      left = left[1:lenL]
      lenL--;
    } else {
      results = append(results, right[0])
      right = right[1:lenR]
      lenR--;
    }
  }

  if lenL > 0 {
    results = append(results, left...)
  } else if lenR > 0 {
    results = append(results, right...)
  }
  return results
}

func mergeSort(values []int, outChan chan []int) (results []int){
  length := len(values)
  middle := length / 2
  if length <= 1 {
    if outChan != nil {
      outChan <- values
    }
    return values
  }
  rightChan := make(chan []int)
  go mergeSort(values[middle:], rightChan)
  left := mergeSort(values[:middle], nil)
  right := <-rightChan
  results = merge(left, right)
  if outChan != nil {
    outChan <- results
  }
  return results
}

func MergeSortP(values []int) {
  numCPU := runtime.NumCPU()
  fmt.Println("usage", numCPU, "CPU")
  runtime.GOMAXPROCS(numCPU)
  outChan := make(chan []int)
  go mergeSort(values, outChan)
  results := <-outChan
  copy(values, results)
}
