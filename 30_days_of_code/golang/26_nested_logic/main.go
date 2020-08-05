package main

import (
    "fmt"
    "strings"
    "time"
    "strconv"
)

func main() {
    var (
        day string
        month string
        year string

        days []string
    )

    for i := 0; i < 2; i++ {
        fmt.Scan(&day)
        fmt.Scan(&month)
        fmt.Scan(&year)

        dString := fmt.Sprintf("%s %s %s", day, month, year)
        days = append(days, dString)
    }

    fmt.Println(CountFine(days[0], days[1]))
}

func CountFine(d1 string, d2 string) (totalValue int) {
    ds := stoI(strings.Split(d1, " "))
    es := stoI(strings.Split(d2, " "))

    a := getDate(ds)
    e := getDate(es)
    diff := a.Sub(e)

    days := int(diff.Hours() / 24)
    months := int(days / 30)

    if days > 364 || a.Year() > e.Year()  {
        totalValue = 10000
        return
    }

    if months > 0 {
        totalValue = months * 500
        return
    }

    if days > 1 {
        totalValue = 15 * days
        return
    }

    return
}

func getDate(d []int) time.Time {
    return time.Date(d[2], time.Month(d[1]), d[0], 0, 0, 0, 0, time.UTC)
}

func stoI(a []string) (r []int) {
    for _, n := range a {
        number, _ := strconv.Atoi(n)
        r = append(r, number)
    }

    return
}
