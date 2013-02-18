package bubblesort

func BubbleSort(values []int) {
  swapped := false

  for i, length := 0, len(values); i < length - 1; i++ {
    swapped = false
    
    for j := 0; j < length - i - 1; j++ {
      if values[j] > values[j + 1] {
        // swap the element
        values[j] , values [j + 1] = values[j + 1], values[j]
        swapped = true
      }
    }

    if !swapped {
      break
    }
  }
}