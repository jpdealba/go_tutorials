package main

import (
	"fmt"
)

func main() {
	// ! ARRAYS

	var intArr [3]int32
	// arrays are contiguous in memory and are indexable
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// initialize array
	var intArr2 [3]int32 = [3]int32{1, 2, 3} // or intArr2 := [3]int32{1, 2, 3} or intArr2 := [...]int32{1,2,3}
	fmt.Println(intArr2)

	// ! SLICES
	// slices are arrays with steroids, it checks if the array has enough room for another number,
	// if not it creates a new array with double the size and copies the old array into the new one
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// append multiple slices
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)

	var intSlice3 []int32 = make([]int32, 3, 5) // make a slice with length 3 and capacity 5, this makes it faster
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice3), cap(intSlice3))

	// ! MAPS
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	// map always return something, if the key is not found it returns the zero value of the value type
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Adam"]
	fmt.Println(age, ok)
	if ok {
		fmt.Println(myMap2["Adam"])
	} else {
		fmt.Println("Not found")
	}

	delete(myMap2, "Adam")

	// ! LOOPS
	// loop a map, this doesnt guarantee the order of the keys
	for name, age := range myMap2 {
		fmt.Printf("%v is %v years old\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index %v has value %v\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
