package main

import (
    "testing"
    "reflect"
)

const errorMessage = "%v was expected but returned %v"

func TestSort(t *testing.T) {
    var (
        input []int64
        expected []int64
        result []int64
        numSwap int
    )

    input = []int64{3, 2, 1}
    expected = []int64{1, 2, 3}
    result, numSwap = Sort(input, 3)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf(errorMessage, expected, result)
    }

    if numSwap != 3 {
        t.Errorf(errorMessage, 3, numSwap)
    }

    input = []int64{1, 2, 3}
    expected = []int64{1, 2, 3}

    result, numSwap = Sort(input, 3)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf(errorMessage, expected, result)
    }

    if numSwap != 0 {
        t.Errorf(errorMessage, 0, numSwap)
    }
}
