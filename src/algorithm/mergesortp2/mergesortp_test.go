package mergesortp2

import "testing"

func TestMergeSortParallel1(t *testing.T) {
  values := []int{5, 4, 3, 2, 1}
  MergeSortParallel(values)

  if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
     values[4] != 5 {
    t.Error("MergeSortParallel() failed. Got", values, "Expected 1 2 3 4 5")
  }
}

func TestMergeSortParallel2(t *testing.T) {
  values := []int{5, 5, 3, 2, 1}
  MergeSortParallel(values)

  if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 ||
     values[4] != 5 {
    t.Error("MergeSortParallel() failed. Got", values, "Expected 1 2 3 5 5")
  }
}

func TestMergeSortParallel3(t *testing.T) {
  values := []int{5}
  MergeSortParallel(values)

  if values[0] != 5 {
    t.Error("MergeSortParallel() failed. Got", values, "Expected 5")
  }
}
