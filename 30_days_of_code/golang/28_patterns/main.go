package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "regexp"
    "sort"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    N := int32(NTemp)

    var names = []string{}

    for NItr := 0; NItr < int(N); NItr++ {
        firstNameEmailID := strings.Split(readLine(reader), " ")

        firstName := firstNameEmailID[0]
        emailID := firstNameEmailID[1]

        if ValidEmail(emailID) {
            names = append(names, firstName)
        }
    }

    SortString(&names)

    for _, name := range names {
        fmt.Println(name)
    }
}

func ValidEmail(address string) bool {
    const email = "@gmail.com"
    r, _ := regexp.Compile(email)
    matched := r.MatchString(address)

    return matched
}

func SortString(names *[]string) {
    sort.Strings(*names)
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

