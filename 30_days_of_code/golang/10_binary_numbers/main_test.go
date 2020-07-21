package main

import (
    "testing"
    "reflect"
)

//

func TestParseBin(t *testing.T) {
    var expected []int32

    expected = []int32{1, 0, 1}

    var result []int32

    result = ParseBin(5)

    if ! reflect.DeepEqual(expected, result) {
        t.Errorf("TestParseBin(5) returned: %v, but %v was expected.", result, expected)
    }

    result = ParseBin(6)

    expected = []int32{0, 1, 1}
    if ! reflect.DeepEqual(expected, result) {
        t.Errorf("TestParseBin(5) returned: %v, but %v was expected.", result, expected)
    }

    result = ParseBin(13)

    expected = []int32{1, 0, 1, 1}
    if ! reflect.DeepEqual(expected, result) {
        t.Errorf("TestParseBin(5) returned: %v, but %v was expected.", result, expected)
    }

    result = ParseBin(143)

    expected = []int32{1, 1, 1 ,1, 0, 0, 0, 1}
    if ! reflect.DeepEqual(expected, result) {
        t.Errorf("TestParseBin(5) returned: %v, but %v was expected.", result, expected)
    }
}

func TestCountConsecutive(t *testing.T) {
    expected := 2

    arr := []int32{1, 1, 0, 0}

    result := CountConsecutive(arr)

    if result != expected {
        t.Errorf("CountConsecutive([]int32) returned: %d | %d was expected.", result, 1)
    }
}
