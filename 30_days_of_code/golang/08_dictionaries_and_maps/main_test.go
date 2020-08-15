package main

import (
    "testing"
    "reflect"
)

func TestSaveContact(t *testing.T) {
    dataSet := []struct{
        input string
        expected  map[string]int
    }{
        {"sam 99912222", map[string]int{"sam": 99912222}},
    }
    agenda := map[string]int{}

    for _, d := range dataSet {
        saveContact(&agenda, d.input)

        if  !reflect.DeepEqual(agenda, d.expected) {
            t.Errorf("%v was expected, but %v was returned", d.expected, agenda)
        }
    }
}

func TestFindContact(t *testing.T) {
    var agenda = map[string]int{"sam": 99912222}
    dataSet := []struct{
        input string
        expected string
    }{
        {"sam", "sam=99912222"},
        {"roge", "Not found"},
    }

    for _, d := range dataSet {
        result := findContact(&agenda, d.input)

        if result != d.expected {
            t.Errorf("%v was expected, but %v was returned", d.expected, result)
        }
    }
}
