package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    var (
        i uint64 = 4
        d float64 = 4.0
        s string = "HackerRank"

        a uint64
        b float64
        c string
    )

    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    j, _ := strconv.ParseInt(scanner.Text(), 10, 64)
    a = uint64(j)

    scanner.Scan()
    b, _ = strconv.ParseFloat(scanner.Text(), 64)

    scanner.Scan()
    c = scanner.Text()

    fmt.Println(SumInt(a, i))
    fmt.Printf("%.1f\n", SumFloat(d, b))
    fmt.Println(Concat(s, c))
}

func SumInt(a, i uint64) uint64 {
    return a + i
}

func SumFloat(a, i float64) float64 {
    n := a + i
    return n
}

func Concat(a, b string) string {
    return fmt.Sprintf("%s %s", a, b)
}
