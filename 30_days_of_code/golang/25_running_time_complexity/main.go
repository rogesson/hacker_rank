package main

import (
    "fmt"
    "bufio"
    "os"
    "io"
    "strconv"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    n, _, err := reader.ReadLine()
    number, _ := strconv.Atoi(string(n))

    if err ==  io.EOF {
        panic("end file reached")
    }

    for i := 0; i < number; i++ {
        newNumber, _, _ := reader.ReadLine()
        number := string(newNumber)
        n, _ := strconv.ParseInt(number, 10, 64)
        if IsPrime(n) {
            fmt.Println("Prime")
        } else {
            fmt.Println("Not prime")
        }
    }
}

func IsPrime(number int64) bool {
    if number == 1 {
        return false
    }

    if number % 1 != 0 {
        return false
    }

    for i := 2; i * i <= int(number); i ++ {
        if int(number) % i == 0 {
            return false
        }
    }

    return number % number == 0
}
