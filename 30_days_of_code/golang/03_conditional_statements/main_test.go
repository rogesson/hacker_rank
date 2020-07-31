package main

import (
    "testing"
)

func TestCheckNumber(t *testing.T) {
    dataset := []struct {
        input int32
        expected string
    }{
        {3, "Weird"},
        {4, "Not Weird"},
        {8, "Weird"},
        {22, "Not Weird"},
        {54, "Not Weird"},
    }

    for _, d := range dataset {
        result := CheckNumber(d.input)

        if result != d.expected {
            t.Errorf("input: %d | %s was expected. result: %s", d.input, d.expected, result)
        }
    }
}
