package main

import (
  "bufio"
  "flag" 
  "fmt"
  "io"
  "os"
  "strconv"
  "time"
  "algorithm/bubblesort"
  "algorithm/qsort"
  "algorithm/mergesort"
  "algorithm/mergesortp"
)

var inFile string 
var outFile string
var algorithm string

func readValues (inFile string) (values []int, err error) {
  file, err := os.Open(inFile)
  if err != nil {
    fmt.Println("Failed to open the input file ", inFile)
    return
  }

  defer file.Close()

  br := bufio.NewReader(file)

  values = make([]int, 0)

  for {
    line, isPrefix, err1 := br.ReadLine()

    if err1 != nil {
      if err1 != io.EOF {
        err = err1
      }
      break
    }

    if isPrefix {
      fmt.Println("A too long line, seems unexpected.")
      return
    }

    str := string(line)

    value, err1 := strconv.Atoi(str)

    if err1 != nil {
      err = err1
      return
    }

    values = append(values, value)
  }
  return values, err
}

func writeValues(values []int, outFile string) error {
  file, err := os.Create(outFile)
  if err != nil {
    fmt.Println("Failed to create the output file ", outFile)
    return err
  }

  defer file.Close()

  for _, value := range values {
    str := strconv.Itoa(value)
    file.WriteString(str + "\n")
  }
  return nil
}

func init(){
  flag.StringVar(&inFile, "i", "unsorted.dat", "File contains values for sorting")
  flag.StringVar(&outFile, "o", "sorted.dat", "File to receive sorted values")
  flag.StringVar(&algorithm, "a", "qsort", "Sort algorithm")
}

func main() {
  flag.Parse()
  fmt.Println("inFile = ", inFile, "outFile = ", outFile, "algorithm = ", algorithm)

  values, err := readValues(inFile)

  if err == nil  {
    startTime := time.Now()
    switch algorithm {
      case "qsort" :
        qsort.QuickSort(values)
      case "bubblesort" :
        bubblesort.BubbleSort(values)
      case "mergesort" :
        mergesort.MergeSort(values)
      case "mergesortp" :
        mergesortp.MergeSortP(values)
      default :{
        fmt.Println("Sorting algorithm", algorithm, "is either unknow or unsupported.")
      }
    }
    endTime := time.Now()
    fmt.Println("The sorting process costs", endTime.Sub(startTime), "to complete")
    writeValues(values, outFile)
  } else {
    fmt.Println(err)
  }
}