package main

import (
  "flag"
  "fmt"
  "os"
  "strconv"
  "time"
  "math/rand"
)

var num int
var outFile string

func init() {
  flag.IntVar(&num, "n", 10000, "Number of element.")
  flag.StringVar(&outFile, "o", "unsorted.dat", "Output file")
}

func writeValues(values []int, outfile string) error {
  file, err := os.Create(outfile)
  if err != nil {
    fmt.Println("Failed to create the output file ", outfile)
    return err
  }

  defer file.Close()

  for _, value := range values {
    str := strconv.Itoa(value)
    file.WriteString(str + "\n")
  }
  return nil
}

func generateValues(num int) (values []int) {
  values = make([]int, num, num)
  for i := 0; i < num; i++ {
    values[i] = rand.Intn(num)
  }
  return values
}

func main() {
  flag.Parse()
  fmt.Println(num, outFile)
  rand.Seed(time.Now().UTC().UnixNano())

  values := generateValues(num)
  writeValues(values, outFile)
}