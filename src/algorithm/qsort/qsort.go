package qsort

const (
  cutoff = 10
)

func partition(values []int, left, right, pivotIndex int) (storeIndex int) {
  pivotValue := values[pivotIndex]
  values[pivotIndex], values[right] = values[right], values[pivotIndex]
  storeIndex = left
  for i := left; i < right; i++ {
    if values[i] < pivotValue {
      values[i], values[storeIndex] = values[storeIndex], values[i]
      storeIndex += 1
    }
  }
  values[storeIndex], values[right] = values[right], values[storeIndex]
  return storeIndex
}

// implement median-of-three partitioning
// ref: 
// http://www.java-tips.org/java-se-tips/java.lang/quick-sort-implementation-with-median-of-three-partitioning-and-cutoff-for-small-a.html
// cutoff value is defined empirically as 10
// ref:
// http://www.cs.wcupa.edu/~rkline/DS/fast-sorts.html
func quickSort(values []int, left, right int) {
  if left < right {
    if right - left < cutoff {
      insertionSort(values, left, right)
      return
    }
    middle := left + (right - left) / 2
    if values[middle] < values[left] {
      values[middle], values[left] = values[left], values[middle]
    }
    if values[right] < values[left] {
      values[right], values[left] = values[left], values[right]
    }
    if values[right] < values[middle] {
      values[right], values[middle] = values[middle], values[right]
    }
    values[middle], values[right - 1] = values[right -1], values[middle]

    pivotNewIndex := partition(values, left, right, right - 1)

    quickSort(values, left, pivotNewIndex - 1)
    quickSort(values, pivotNewIndex + 1, right) 
  }
}

func insertionSort(values []int, left, right int) {
  for i := left + 1; i <= right; i++ {
    checkValue := values[i]
    j := i
    for ; j > left && checkValue < values[j - 1]; j-- {
      values[j] = values[j - 1]
    }
    values[j] = checkValue
  }
}

func QuickSort(values []int) {
  quickSort(values, 0, len(values) - 1)
}