package main

import(
    "testing"
)

func TestSumInt(t *testing.T){
    sum := SumInt(1, 3)

    if sum != 4 {
        t.Errorf("%d was expected, but %d was returned", 4, sum)
    }
}

func TestSumFloat(t *testing.T) {
    sum := SumFloat(4.0, 4.0)

    if sum != 8.0 {
        t.Errorf("%f was expected, but %f was returned", 8.0, sum)
    }
}

func TestConcat(t *testing.T) {
    text := "HackerRank is the best place to learn and practice coding!"
    str := Concat("HackerRank", "is the best place to learn and practice coding!")

    if str != text {
        t.Errorf("%s was expected, but %s was returned", text, str)
    }
}
