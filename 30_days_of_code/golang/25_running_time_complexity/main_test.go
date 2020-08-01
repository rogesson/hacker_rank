package main

import (
    "testing"
)

func TestIsPrime(t *testing.T) {
    dataSet := []struct {
        input int64
        result bool
    }{
        {1, false},
        {2, true},
        {12, false},
        {5, true},
        {7, true},
    }


    for _, d := range dataSet {
        prime := IsPrime(d.input)

        if prime != d.result {
            t.Errorf("Input %d shoud be %v, but was, %v\n", d.input, d.result, prime)
        }
    }
}
