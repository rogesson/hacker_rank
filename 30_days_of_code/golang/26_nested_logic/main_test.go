package main

import (
    "testing"
)

func TestCountFine(t *testing.T) {
    dataSet := []struct{
        d1 string
        d2 string
        result int
    }{
        {"9 6 2015", "6 6 2015", 45},
        {"8 4 12", "1 1 1", 10000},
        {"1 1 2010", "31 12 2009", 10000},
        {"31 8 2004", "20 1 2004", 3500},
    }

    for _, d := range dataSet {
        r := CountFine(d.d1, d.d2)
        if r != d.result {
            t.Errorf("%v was expecetd, but %v was returned", d.result, r)
        }
    }
}
