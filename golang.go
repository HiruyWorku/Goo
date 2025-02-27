// Go Programming Basics: Structs, Maps, and Iterations

package main

import "fmt"

// Defining a struct Person with fields name, age, job, and salary
type Person struct {
    name   string
    age    int
    job    string
    salary int
}

// Defining a struct Student with fields name, year, major, and gpa
type Student struct {
    name  string
    year  string
    major string
    gpa   float64
}

// Function to print person details
func printPerson(pers Person) {
    fmt.Println("Name: ", pers.name)
    fmt.Println("Age: ", pers.age)
    fmt.Println("Job: ", pers.job)
    fmt.Println("Salary: ", pers.salary)
}

func main() {
    var pers1 Person
    var student1 Student

    // Assigning values to Person struct
    pers1.name = "Hege"
    pers1.age = 45
    pers1.job = "Teacher"
    pers1.salary = 6000

    // Assigning values to Student struct
    student1.name = "Hiruy"
    student1.year = "Sophomore"
    student1.major = "Computer Science"
    student1.gpa = 4.0

    // Printing struct values
    fmt.Println("Name: ", pers1.name)
    fmt.Println("Age: ", pers1.age)
    fmt.Println("Job: ", pers1.job)
    fmt.Println("Salary: ", pers1.salary)

    fmt.Println("Name: ", student1.name)
    fmt.Println("Year: ", student1.year)
    fmt.Println("Major: ", student1.major)
    fmt.Println("GPA: ", student1.gpa)

    // Calling function to print person details
    printPerson(pers1)

    // Creating and printing a map
    var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
    fmt.Printf("a\t%v\n", a)
    fmt.Printf("b\t%v\n", b)

    // Using make() to create maps
    var c = make(map[string]string)
    c["brand"] = "Ford"
    c["model"] = "Mustang"
    c["year"] = "1964"

    d := make(map[string]int)
    d["Oslo"] = 1
    d["Bergen"] = 2
    d["Trondheim"] = 3
    d["Stavanger"] = 4

    fmt.Printf("c\t%v\n", c)
    fmt.Printf("d\t%v\n", d)

    // Checking if maps are nil
    var e = make(map[string]string)
    var f map[string]string
    fmt.Println(e == nil) // Output: false
    fmt.Println(f == nil) // Output: true

    // Updating and adding elements in a map
    var g = make(map[string]string)
    g["brand"] = "Ford"
    g["model"] = "Mustang"
    g["year"] = "1964"
    fmt.Println(g)

    g["year"] = "1970"
    g["color"] = "red"
    g["laptop-brand"] = "HP"
    fmt.Println(g)

    // Maps as references
    var h = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
    i := h
    fmt.Println(h)
    fmt.Println(i)

    i["year"] = "1970"
    fmt.Println("After change to i:")
    fmt.Println(h)
    fmt.Println(i)

    // Iterating over a map
    j := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
    for k, v := range j {
        fmt.Printf("%v : %v, ", k, v)
    }
    fmt.Println()

    var s []string
    s = append(s, "one", "two", "three", "four")
    for _, element := range s {
        fmt.Printf("%v : %v, ", element, j[element])
    }
    fmt.Println()

    // Deleting from map
    delete(g, "brand")
    fmt.Println(g)

    // Checking for key existence
    val1, ok1 := g["color"]
    val2, ok2 := g["road"]
    val3, ok3 := g["laptop-brand"]
    _, ok4 := g["model"]
    fmt.Println(val1, ok1) // Output: red true
    fmt.Println(val2, ok2) // Output: false
    fmt.Println(val3, ok3) // Output: HP true
    fmt.Println(ok4)       // Output: true
}
