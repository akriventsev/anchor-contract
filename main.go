package main

import "fmt"

// Simple function that adds two numbers
func add(a, b int) int {
    return a + b
}

// Simple function that greets a person
func greet(name string) string {
    return "Hello, " + name + "!"
}

// Simple function that checks if a number is even
func isEven(num int) bool {
    return num%2 == 0
}

func main() {
    // Test the add function
    result := add(5, 3)
    fmt.Printf("5 + 3 = %d\n", result)
    
    // Test the greet function
    message := greet("World")
    fmt.Println(message)
    
    // Test the isEven function
    number := 7
    if isEven(number) {
        fmt.Printf("%d is even\n", number)
    } else {
        fmt.Printf("%d is odd\n", number)
    }
}