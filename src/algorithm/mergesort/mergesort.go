package mergesort

func merge(left, right []int) (results []int) {
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

func mergeSort(values []int) (results []int){
  length := len(values)
  middle := length / 2
  if length <= 1 {
    return values
  }
  left := mergeSort(values[:middle])
  right := mergeSort(values[middle:])
  results = merge(left, right)
  return results
}

func MergeSort(values []int) {
  results := mergeSort(values)
  copy(values, results)
}