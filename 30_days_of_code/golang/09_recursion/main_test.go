package main

import (
    "testing"
)

func TestFactorial(t *testing.T) {
    dataSet := []struct{
        input int
        expected int
    }{
        {3, 6},
    }

    for _, d := range dataSet {
        result := factorial(d.input)

        if result != d.expected {
            t.Errorf("want %v got %v", d.expected, result)
        }
    }
}
