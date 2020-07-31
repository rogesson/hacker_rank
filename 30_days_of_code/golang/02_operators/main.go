package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func Solve(meal_cost float64, tip_percent int32, tax_percent int32) float64 {
    p := meal_cost * float64(tip_percent) / 100

    tax := meal_cost * float64(tax_percent) / 100
    meal_cost += p

    fmt.Printf("%2.f\n", meal_cost + tax)
    result := fmt.Sprintf("%2.f", meal_cost + tax)
    totalCost, _ := strconv.ParseFloat(result, 64)
    return totalCost
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    Solve(meal_cost, tip_percent, tax_percent)
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

