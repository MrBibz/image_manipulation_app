package main

import "fmt"

func TestIf(value int) {
	if value > 0 {
		fmt.Println(fmt.Sprintf("%d is positive", value))
	} else if value < 0 {
		fmt.Println(fmt.Sprintf("%d is negative", value))
	} else {
		fmt.Println(fmt.Sprintf("%d is zero", value))
	}
}

func TestSwitch(value int) {
	switch {
	case value > 0:
		fmt.Println(fmt.Sprintf("Value %d is positive", value))
	case value < 0:
		fmt.Println(fmt.Sprintf("Value %d is negative", value))
	default:
		fmt.Println(fmt.Sprintf("Value %d is zero", value))
	}
}

func TestFor() {
	fmt.Print("Counting from 1 to 10 using a for loop :\n")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func TestWhile() {
	fmt.Print("Counting from 1 to 10 using a while loop :\n")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func TestList() {
	fmt.Println("Creating a list of 5 elements :")
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(list)

	fmt.Println("Adding an element to the list :")
	list = append(list, 6)
	fmt.Println(list)

	fmt.Println("Removing the last element of the list :")
	list = list[:len(list)-1]
	fmt.Println(list)
}

func TestMap() {
	fmt.Println("Creating a map of 3 elements :")
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	fmt.Println("Adding an element to the map :")
	m["four"] = 4
	fmt.Println(m)

	fmt.Println("Removing an element from the map :")
	delete(m, "two")
	fmt.Println(m)
}

func main() {
	// Tests conditional statements
	fmt.Println("\nTest if statement:")
	TestIf(5)
	TestIf(-5)
	TestIf(0)

	fmt.Println("\nTest switch statement:")
	TestSwitch(5)
	TestSwitch(-5)
	TestSwitch(0)

	// Test iteration statements
	fmt.Println("\nTest for statement:")
	TestFor()

	fmt.Println("\nTest while statement:")
	TestWhile()

	// Test collections
	fmt.Println("\nTest list:")
	TestList()

	fmt.Println("\nTest map:")
	TestMap()

	// Test functions
	fmt.Println("\nTest functions:")
	fmt.Println("\nEvery test is in a function in this file.")
}
