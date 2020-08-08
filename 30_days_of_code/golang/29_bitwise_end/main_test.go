package main

import (
    "testing"
    "reflect"
)

func TestAllPossibleValues(t *testing.T) {
    input := []Number{
        Number{5, 2},
        Number{8, 5},
        Number{2, 2},
    }

    expected := []int{1, 4, 0}

    result := AllPossibleValues(&input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("%v is expected but %v was returned", expected, result)
    }
}
