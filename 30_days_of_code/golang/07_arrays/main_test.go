package main

import (
    "testing"
    "reflect"
)

func TestReverseArray(t *testing.T) {
    input := []int32{1, 4, 3, 2}
    expected := []int32{2, 3, 4, 1}

    result := reverseArray(&input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("%v is expected, but %v was returned", expected, result)
    }
}
