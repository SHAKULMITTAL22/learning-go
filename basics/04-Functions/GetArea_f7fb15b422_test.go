package main

import "testing"

func TestGetArea_f7fb15b422(t *testing.T) {
    // Test case 1: Valid input
    l := 5
    b := 10
    expectedArea := 50
    actualArea := getArea(l, b)
    if actualArea != expectedArea {
        t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
    }

    // Test case 2: Invalid input (negative values)
    l = -5
    b = -10
    expectedArea := 0
    actualArea = getArea(l, b)
    if actualArea != expectedArea {
        t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
    }

    // Test case 3: Invalid input (zero values)
    l = 0
    b = 0
    expectedArea := 0
    actualArea = getArea(l, b)
    if actualArea != expectedArea {
        t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
    }

    // Test case 4: Edge case (large values)
    l = 1000000000
    b = 1000000000
    expectedArea := 1000000000000000000
    actualArea = getArea(l, b)
    if actualArea != expectedArea {
        t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
    }

    // Test case 5: Edge case (small values)
    l = 1
    b = 1
    expectedArea = 1
    actualArea = getArea(l, b)
    if actualArea != expectedArea {
        t.Errorf("Expected area %d, got %d", expectedArea, actualArea)
    }
}
