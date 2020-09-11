package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	//  Declare and Create Slice - two steps
	var notes []string
	notes = make([]string, 7)

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	notes[3] = "fa"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[2])

	// Declare and Create in 1 step of a slice
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	fmt.Println("-----------------------------------")
	fmt.Println("Len of notes  =", len(notes))
	fmt.Println("Len of primes =", len(primes))

	fmt.Println("-----------------------------------")
	fmt.Println("2 for loops ")
	// slice literals
	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, val := range letters {
		fmt.Println(val)
	}

	fmt.Println("-----------------------------------")
	fmt.Println("Create a slice 1 from an array")
	// Base Array
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// Create Slice from array
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)
	fmt.Println("-----------------------------------")
	fmt.Println("Create a slice 2 from an array")
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)

	fmt.Println("-----------------------------------")
	fmt.Println("Create a slice 3 from an array")
	i, j = 2, 5
	slice3 := underlyingArray[i:j]
	fmt.Println(slice3)

	//fmt.Println("-----------------------------------")
	//fmt.Println("Create a slice 4 from an array")
	//i, j = 2, 6
	//slice4 := underlyingArray[i:j]
	//fmt.Println(slice4)

	fmt.Println("-----------------------------------")
	fmt.Println("Create a slice 5 from an array")
	i, j = 2, 5
	// [i:] to the end of the array
	slice5 := underlyingArray[i:]
	fmt.Println(slice5)

	fmt.Println("-----------------------------------")
	fmt.Println("Create a slice 6 from an array")
	slice6 := []string{"a", "b"}
	fmt.Println(slice6, "len =", len(slice6))
	slice6 = append(slice6, "c")
	fmt.Println(slice6, "len =", len(slice6))
	slice6 = append(slice6, "d", "e")
	fmt.Println(slice6, "len =", len(slice6))

	fmt.Println("-----------------------------------")
	fmt.Println("Create a int slice and string slice - empty un-created ")
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, string slice %#v\n\n", intSlice, stringSlice)

}
