package main

import (
    "testing"
)

func TestPrintMessages(t *testing.T) {
    var (
        expected bool
        result bool
        messages []string
    )

    messages = []string{"Hello, roger", "some text"}
    result = PrintMessages(messages)
    expected = true

    if  expected != result {
        t.Errorf("TestPrintText([]string) returned %v, %v was expected", result, expected)
    }
}
