package main

import "fmt"

// Function with multiple return values
func divide(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    quotient := dividend / divisor
    return quotient, nil
}