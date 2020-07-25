package main

import (
    "testing"
    "reflect"
)

func TestCreateHourglasses(t *testing.T) {
    var input [][]int32
    var expected [][]int32
    var result [][]int32

    expected = [][]int32 {
        {1, 1, 1, 1, 1, 1, 1},
        {1, 1, 0, 0, 1, 1, 0},
        {1, 0, 0, 0, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 0},
        {0, 1, 0, 1, 0, 0, 2},
        {1, 0, 0, 1, 0, 2, 4},
        {0, 0, 0, 0, 2, 4, 4},
        {0, 0, 0, 0, 4, 4, 0},
        {1, 1, 1, 0, 0, 0, 0},
        {1, 1, 0, 2, 0, 0, 2},
        {1, 0, 0, 4, 0, 2, 0},
        {0, 0, 0, 4, 2, 0, 0},
        {0, 0, 2, 0, 0, 0, 1},
        {0, 2, 4, 0, 0, 1, 2},
        {2, 4, 4, 2, 1, 2, 4},
        {4, 4, 0, 0, 2, 4, 0},
    }

    input = [][]int32 {
        {1, 1, 1, 0, 0, 0},
        {0, 1, 0, 0, 0, 0},
        {1, 1, 1, 0, 0, 0},
        {0, 0, 2, 4, 4, 0},
        {0, 0, 0, 2, 0, 0},
        {0, 0, 1, 2, 4, 0},
    }

    result = CreateHourglasses(input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("CountConsecutive([]int32) returned: %d | %d was expected.", result, expected)
    }

    input = [][]int32 {
        {-1, -1, 0,  -9, -2, -2},
        {-2, -1, -6, -8, -2, -5},
        {-1, -1, -1, -2, -3, -4},
        {-1, -9, -2, -4, -4, -5},
        {-7, -3, -3, -2, -9, -9},
        {-1, -3, -1, -2, -4, -5},
    }

    expected = [][]int32 {
        {-1, -1,  0, -1, -1, -1, -1},
        {-1,  0, -9, -6, -1, -1, -2},
        { 0, -9, -2, -8, -1, -2, -3},
        {-9, -2, -2, -2, -2, -3, -4},
        {-2, -1, -6, -1, -1, -9, -2},
        {-1, -6, -8, -1, -9, -2, -4},
        {-6, -8, -2, -2, -2, -4, -4},
        {-8, -2, -5, -3, -4, -4, -5},
        {-1, -1, -1, -9, -7, -3, -3},
        {-1, -1, -2, -2, -3, -3, -2},
        {-1, -2, -3, -4, -3, -2, -9},
        {-2, -3, -4, -4, -2, -9, -9},
        {-1, -9, -2, -3, -1, -3, -1},
        {-9, -2, -4, -3, -3, -1, -2},
        {-2, -4, -4, -2, -1, -2, -4},
        {-4, -4, -5, -9, -2, -4, -5},
    }

    result = CreateHourglasses(input)

    if !reflect.DeepEqual(result, expected) {
        t.Errorf("CountConsecutive([]int32) returned: %d | %d was expected.", result, expected)
    }
}

func TestBiggerHourglass(t *testing.T) {
    var hourglasses [][]int32
    var result int32
    var expected int32

    hourglasses = [][]int32 {
        {1, 1, 1, 1, 1, 1, 1},
        {1, 1, 0, 0, 1, 1, 0},
        {1, 0, 0, 0, 1, 0, 0},
        {0, 0, 0, 0, 0, 0, 0},
        {0, 1, 0, 1, 0, 0, 2},
        {1, 0, 0, 1, 0, 2, 4},
        {0, 0, 0, 0, 2, 4, 4},
        {0, 0, 0, 0, 4, 4, 0},
        {1, 1, 1, 0, 0, 0, 0},
        {1, 1, 0, 2, 0, 0, 2},
        {1, 0, 0, 4, 0, 2, 0},
        {0, 0, 0, 4, 2, 0, 0},
        {0, 0, 2, 0, 0, 0, 1},
        {0, 2, 4, 0, 0, 1, 2},
        {2, 4, 4, 2, 1, 2, 4},
        {4, 4, 0, 0, 2, 4, 0},
    }

    expected = 19
    result = BiggerHourglass(hourglasses)

    if result != expected {
        t.Errorf("CountConsecutive([]int32) returned: %d | %d was expected.", result, expected)
    }

    hourglasses = [][]int32 {
        {-1, -1,  0, -1, -1, -1, -1},
        {-1,  0, -9, -6, -1, -1, -2},
        { 0, -9, -2, -8, -1, -2, -3},
        {-9, -2, -2, -2, -2, -3, -4},
        {-2, -1, -6, -1, -1, -9, -2},
        {-1, -6, -8, -1, -9, -2, -4},
        {-6, -8, -2, -2, -2, -4, -4},
        {-8, -2, -5, -3, -4, -4, -5},
        {-1, -1, -1, -9, -7, -3, -3},
        {-1, -1, -2, -2, -3, -3, -2},
        {-1, -2, -3, -4, -3, -2, -9},
        {-2, -3, -4, -4, -2, -9, -9},
        {-1, -9, -2, -3, -1, -3, -1},
        {-9, -2, -4, -3, -3, -1, -2},
        {-2, -4, -4, -2, -1, -2, -4},
        {-4, -4, -5, -9, -2, -4, -5},
    }

    expected = -6
    result = BiggerHourglass(hourglasses)

    if result != expected {
        t.Errorf("CountConsecutive([]int32) returned: %d | %d was expected.", result, expected)
    }
}
