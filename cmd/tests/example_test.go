package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := add(2, 3)
    assert.Equal(t, 5, result, "Сумма должна быть равна 5")
}

func add(a, b int) int {
    return a + b
}
