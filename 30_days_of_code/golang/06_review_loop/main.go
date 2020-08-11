package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    count, _ := strconv.Atoi(readLine(reader))

    for i := 0; i < count; i++ {
        word := readLine(reader)
        fmt.Println(parseWord(word))
    }
}

func parseWord(word string) string {
    var evenWords []string
    var oddWords []string

    for i, w := range strings.Split(word, "") {
        if i % 2 == 0 {
            evenWords = append(evenWords, w)
        } else {
            oddWords = append(oddWords, w)
        }
    }

    strEvenWords := strings.Join(evenWords, "")
    strOddWords := strings.Join(oddWords, "")

    return fmt.Sprintf("%s %s", strEvenWords, strOddWords)
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


