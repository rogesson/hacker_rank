package main

import (
    "fmt"
    "bufio"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    size := readLine(reader)
    str, _, _ := reader.ReadLine()
    array := strings.Split(string(str), " ")

    sortedArray, numSwap := Sort(StringsToInt(array), int(size))

    fmt.Printf("Array is sorted in %d swaps.\n", numSwap)
    fmt.Printf("First Element: %d\n", sortedArray[0])
    fmt.Printf("Last Element: %d\n", sortedArray[size -1])
}

func Sort(array []int64, size int) ([]int64, int) {
    numSwap := 0

    for i := 0; i < size; i++ {
        for j := 0; j < size - 1; j++ {
            if array[j] > array[j + 1] {
                aux := array[j + 1]
                array[j + 1] = array[j]
                array[j] = aux
                numSwap++
            }
        }
    }

    return array, numSwap
}

func StringsToInt(a []string) (newArr []int64) {
    for _, element := range a {
        num, _ := strconv.Atoi(element)
        newArr = append(newArr, int64(num))
    }

    return
}

func readLine(reader *bufio.Reader) (num int64) {
    str, _, err := reader.ReadLine()

    if (err == io.EOF) {
        panic("end of file")
    }

    num, _ = strconv.ParseInt(string(str), 10, 64)

    return
}
