package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    "strconv"
)

func factorial(n int32) int32 {
    if n < 1 {
        return 1
    }

    return n * factorial(n - 1)
}

func main() {
    scanner := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    str, _, err := scanner.ReadLine()

    if err == io.EOF {
       panic(err)
    }

    number, _ := strconv.Atoi(string(str))

    f := factorial(int32(number))

    fmt.Println(f)
}
