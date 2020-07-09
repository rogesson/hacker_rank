package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
  "strings"
)

func main() {
  reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

  NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
  checkError(err)
  N := int32(NTemp)

  if N % 2 != 0 {
    fmt.Println("Weird")
  } else {
    switch number := N; {
    case number >= 2 && number <= 5:
      fmt.Println("Not Weird")
    case number >= 6 && number <= 20:
      fmt.Println("Weird")
    case number > 20:
      fmt.Println("Not Weird")
    default:
      fmt.Println("Not Weird")
    }
  }
}

func readLine(reader *bufio.Reader) string {
  str, _, err := reader.ReadLine()
  if err == io.EOF {
    return ""
  }

  return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
  if err != nil {
    panic(err)
  }
}

