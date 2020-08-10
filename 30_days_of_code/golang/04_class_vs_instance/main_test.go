package main

import (
    "testing"
)

const errorMessage string = "%v is expected, but %v was returned"

func TestNewPerson(t *testing.T) {
    dataSet := []struct{
        input int
        expected int
    }{
        {18, 18},
        {-5, 0},
    }

    for _, d := range dataSet {
        p := person{d.input}

        result := p.NewPerson(d.input).age

        if d.expected != result {
            t.Errorf(errorMessage, d.expected, result)
        }
    }
}

func TestAmIOld(t *testing.T) {
    dataSet := []struct{
        input int
        expected string
    }{
        {10, "You are young."},
        {13, "You are a teenager."},
        {17, "You are a teenager."},
        {18, "You are old."},
        {19, "You are old."},
    }

    for _, d := range dataSet {

        p := person{age: d.input}
        result := p.amIOld()

        if d.expected != result {
            t.Errorf(errorMessage, d.expected, result)
        }
    }
}
