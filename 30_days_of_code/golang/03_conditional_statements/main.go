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

    fmt.Println(CheckNumber(N))
}

func CheckNumber(n int32) (text string) {
    if n % 2 != 0 {
        text = "Weird"
    } else {
        switch number := n; {
        case number >= 2 && number <= 5:
            text = "Not Weird"
        case number >= 6 && number <= 20:
            text = "Weird"
        case number > 20:
            text = "Not Weird"
        default:
            text = "Not Weird"
        }
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

