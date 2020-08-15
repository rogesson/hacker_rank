package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int(nTemp)

    agenda := make(map[string]int)

    for i := 0; i < n; i++ {
        contact := readLine(reader)

        saveContact(&agenda, contact)
    }

    for i := n;; i++ {
        name := readLine(reader)

        if name == "" {
            break
        }

        contact := findContact(&agenda, name)
        fmt.Println(contact)
    }
}

func saveContact(agenda *map[string]int, contact string) {
    c := strings.Split(contact, " ")
    name := c[0]
    number, _ := strconv.Atoi(c[1])

    (*agenda)[name] = number
}

func findContact(agenda *map[string]int, name string) string {
    number := (*agenda)[name]

    if number == 0 {
        return "Not found"
    }

    return fmt.Sprintf("%s=%d", name, number)
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

