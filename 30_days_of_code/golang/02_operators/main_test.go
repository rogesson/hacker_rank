package main

import (
    "testing"
)

func TestSolve(t *testing.T) {
    var mealCost float64 = 12.00
    var tipPercent int32 = 20
    var taxPercent int32 = 8

    var expected float64 = 15

    result := Solve(mealCost, tipPercent, taxPercent)

    if result != expected {
        t.Errorf("%f is expected, but %f was returned", expected, result)
    }
}
