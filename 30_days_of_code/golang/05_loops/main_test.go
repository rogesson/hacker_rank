package main

import (
    "testing"
)

func TestCalculate(t *testing.T) {
    dataSet := []struct{
        input, index, expected int32
    }{
        {2, 1, 2},
        {2, 2, 4},
        {2, 3, 6},
    }

    for _, d := range dataSet {
        result := calculate(d.input, d.index)

        if d.expected != result {
            t.Errorf("%v was expected, but %v was returned", d.expected, result)
        }
    }
}
