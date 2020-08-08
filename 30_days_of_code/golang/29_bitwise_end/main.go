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

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)
    var numbers = []Number{}

    for tItr := 0; tItr < int(t); tItr++ {
        nk := strings.Split(readLine(reader), " ")

        nTemp, err := strconv.ParseInt(nk[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(nk[1], 10, 64)
        checkError(err)
        k := int32(kTemp)

        numbers = append(numbers, Number{n, k})
    }

    values := AllPossibleValues(&numbers)

    for _, v := range values {
        fmt.Println(v)
    }
}

type Number struct {
    a int32
    b int32
}

func AllPossibleValues(numbers *[]Number) []int {
    var r = []int{}

    for _, num := range *numbers {
        n := int(num.a)
        k := num.b
        var x int

        for i := 0; i < n - 1; i ++ {
            for j := 1; j < n - i; j++ {
                a :=  i + 1
                b := j + i + 1
                ab := a & b
                if ab < int(k) {
                    if ab > x {
                        x = ab
                    }
                }
            }
        }

        r = append(r, x)
    }

    return r
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

