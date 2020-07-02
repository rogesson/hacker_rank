package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank"
  
  var a uint64
  var b float64
  var c string

  scanner := bufio.NewScanner(os.Stdin)
  
  scanner.Scan()
  j, _ := strconv.ParseInt(scanner.Text(), 10, 64)
  a = uint64(j)
  
  scanner.Scan()
  b, _ = strconv.ParseFloat(scanner.Text(), 64)

  scanner.Scan()
  c = scanner.Text()

  fmt.Println(a + i)
  fmt.Printf("%.1f\n", float64(d + b))
  fmt.Printf("%s %s\n", s, c)
}
