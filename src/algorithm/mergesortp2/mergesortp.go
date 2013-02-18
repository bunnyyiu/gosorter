package mergesortp2


/**
  This is a parallel merge sort presented in 
  Introdction to Algorithm rev 3 Section 27.3 page 799 to 803.

  This one should perform better than the traditional mergesort (mergesort) and 
  parallel mergesort (mergesortp). But in our expierment, it is 100 times slower
  than the serial one (mergesort).
**/

import (
  "os"
  "fmt"
)

func maxInt (x, y int) (result int) {
  if x >= y {
    result = x
  } else {
    result = y
  }
  return result
}

func binarySearch(x int, values []int, left, right int) int{
  low := left
  high := maxInt(left, right + 1)
  for ; low < high; {
    mid := (low + high) / 2
    if x <= values[mid] {
      high = mid
    } else {
      low = mid + 1    
    }
  }
  return high
}

func mergeParallel(values []int, p1, r1, p2, r2 int, results []int, p3 int,
                   outChan chan int) {
  n1 := r1 - p1 + 1
  n2 := r2 - p2 + 1

  if n1 < n2 {
    p1, p2 = p2, p1
    r1, r2 = r2, r1
    n1, n2 = n2, n1
  }
  if n1 == 0 {
    if outChan != nil {
      outChan <- 1
    }
    return
  }
  q1 := (p1 + r1) / 2
  q2 := binarySearch(values[q1], values, p2, r2)
  q3 := p3 + (q1 - p1) + (q2 - p2)
  results[q3] = values[q1]
  childChan := make(chan int)
  go mergeParallel(values, p1, q1 - 1, p2, q2 - 1, results, p3, childChan)
  mergeParallel(values, q1 + 1, r1, q2, r2, results, q3 + 1, nil)
  <-childChan

  if outChan != nil {
    outChan <- 1
  }
  return
}

func mergeSortParallel(values []int, p, r int, results []int, s int,
                       outChan chan int) {
  n := r - p + 1
  if n == 1 {
    results[s] = values[p]
  } else {
    T := make([]int, n, n)
    q := (p + r) / 2
    q2 := q - p

    childChan := make(chan int)
    go mergeSortParallel(values, p, q, T, 0, childChan)
    mergeSortParallel(values, q + 1, r, T, q2 + 1, nil)
    <-childChan
    mergeParallel(T, 0, q2, q2 + 1, n - 1, results, s, nil)
  }
  if outChan != nil {
    outChan <- 1
  }
}

func MergeSortParallel(values []int) {
  fmt.Println("using", os.Getenv("GOMAXPROCS"), "CPUs")
  n := len(values)
  results := make([]int, n, n)
  mergeSortParallel(values, 0, n - 1, results, 0, nil)
  copy(values, results)
}
