// Go Programming Basics: Syntax, Comments, Variables, Constants, Output, and Data Types

// Package declaration
package main

// Importing required packages
import (
    "fmt"
    "math"
)

// Single-line comment: This is a simple Go program
/*
Multi-line comment:
This program demonstrates basic Go syntax,
including variables, constants, output, and data types.
*/

// Defining constants
const Pi = 3.14159

func main() {
    // Declaring and initializing variables
    var name string = "John"
    var age int = 25
    var height float64 = 5.9
    var isStudent bool = false

    // Short variable declaration (type inferred)
    country := "USA"
    score := 95.5

    // Printing output
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Height:", height)
    fmt.Println("Is Student:", isStudent)
    fmt.Println("Country:", country)
    fmt.Println("Score:", score)
    fmt.Println("Value of Pi:", Pi)

    // Data types demonstration
    var intVar int = 10
    var floatVar float64 = 20.5
    var stringVar string = "Hello, Go!"
    var boolVar bool = true

    fmt.Println("Integer Variable:", intVar)
    fmt.Println("Float Variable:", floatVar)
    fmt.Println("String Variable:", stringVar)
    fmt.Println("Boolean Variable:", boolVar)

    // Using the math package
    fmt.Println("Square root of 16:", math.Sqrt(16))
}
