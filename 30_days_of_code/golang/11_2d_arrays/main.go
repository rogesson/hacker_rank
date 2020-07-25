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

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    hourglasses := CreateHourglasses(arr)
    maxNumber := BiggerHourglass(hourglasses)
    fmt.Println(maxNumber)
}

func CreateHourglasses(h [][]int32) (hourglasses [][]int32) {
    for x := 0; x < 4; x++ {
        for y := 0; y < 4; y ++ {
            var hourglass []int32

            for i := 0; i < 3; i++ {
                for j := 0; j < 3; j++ {
                    if i == 1 && j == 0 || i == 1 && j == 2 {
                        continue
                    }

                    newX := i + x
                    newY := j + y

                    hourglass = append(hourglass, h[newX][newY])
                }
            }
            hourglasses = append(hourglasses, hourglass)
        }
    }
    return
}

func BiggerHourglass(hourglasses [][]int32) int32 {
    var counts []int32

    for i := 0; i < len(hourglasses); i++ {
        var count int32

        for _, h := range hourglasses[i] {
            count += h
        }
        counts = append(counts, count)
    }

    var max int32 = counts[0]
    for _, c := range counts {
        if c > max {
            max = c
        }
    }

    return max
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

