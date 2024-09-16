package main

import (
    "testing"
)

func TestPartialMultiplication_1dcd61c0ca(t *testing.T) {
    // Test case 1: Multiply by 2
    expected := 4
    actual := partialMultiplication(2)(2)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }

    // Test case 2: Multiply by 0
    expected = 0
    actual = partialMultiplication(0)(2)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }

    // Test case 3: Multiply by -1
    expected = -2
    actual = partialMultiplication(-1)(2)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }

    // Test case 4: Multiply by a large number
    expected = 1000000
    actual = partialMultiplication(1000)(1000)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }

    // Test case 5: Multiply by a negative large number
    expected = -1000000
    actual = partialMultiplication(-1000)(1000)
    if actual != expected {
        t.Errorf("Expected %d, got %d", expected, actual)
    }
}
