package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    "strconv"
)

func factorial(n int) int {
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

    f := factorial(number)

    fmt.Println(f)
}
