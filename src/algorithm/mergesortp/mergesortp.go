package mergesortp

/**
  According to Introduction to Algorithm rev 3
  This algorithm performance should almost the same as the serial one.

  But in our experiemt, this one is slower than the serial one 10 times more if
  GOMAXPROCS = 2.

  ref Introduction to Algorithm, page 798
*/

import (
  "os"
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

func mergeSortParallel(values []int, readyChan chan int) {
  defer func () {
    if readyChan != nil {
      readyChan <- 1
    }
  }()

  length := len(values)
  middle := length / 2
  if length <= 1 {
    return
  }
  syncChan := make(chan int)
  left := values[:middle]
  right := values[middle:]
  go mergeSortParallel(left, syncChan)
  mergeSortParallel(right, nil)
  <-syncChan
  results := merge(left, right)
  copy(values, results)
  return
}

func MergeSortParallel(values []int) {
  fmt.Println("using", os.Getenv("GOMAXPROCS"), "CPUs")
  mergeSortParallel(values, nil)
}
