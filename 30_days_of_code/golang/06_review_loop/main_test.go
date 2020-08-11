package main

import (
    "testing"
)

func TestParseWord(t *testing.T) {
    dataSet := []struct{
        input, expected string
    }{
        {"Hacker", "Hce akr"},
        {"Rank", "Rn ak"},
    }

    for _, d := range dataSet {
        result := parseWord(d.input)

        if result != d.expected {
            t.Errorf("%v was expected, but %v as returned", d.expected, result)
        }
    }
}
