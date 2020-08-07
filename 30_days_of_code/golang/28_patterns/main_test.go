package main

import (
    "testing"
    "reflect"
)

func TestValidNames(t *testing.T) {
    var input = []struct{
        address string
        result bool
    }{
        {"rogesson@hotmail.com", false},
        {"rogessonb@gmail.com", true},
    }

    for _, data := range input {
        result := ValidEmail(data.address)

        if result != data.result {
            t.Errorf("%s | %v was expected, but %v was returned", data.address, data.result, result)
        }
    }
}

func TestSortString(t *testing.T) {
    var input = []string{"riya", "julia", "julia", "samantha", "tanya"}
    var expected = []string{"julia", "julia", "riya", "samantha", "tanya"}

    SortString(&input)
    result := input

    if !reflect.DeepEqual(expected, result) {
        t.Errorf("%v was expected, but %v was returned", expected, result)
    }
}
