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

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)
    binary := ParseBin(n)
    fmt.Println(CountConsecutive(binary))
}

func CountConsecutive(binary []int32) (max_counter int) {
    counter := 0
    max_counter = 1

    for i := 0; i < len(binary); i++ {
        remainder := binary[i]
        if remainder == 1 {
            counter += 1
            if counter > max_counter {
                max_counter = counter
            }
        } else {
            counter = 0
        }
    }

    return
}

func ParseBin(n int32) (binary []int32) {
    for i := n; i > 0; {
        remainder := n % 2
        n /= 2
        i = n

        binary = append(binary, remainder)
    }

    return
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

