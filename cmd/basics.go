package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

func basicGolang() {
	fmt.Println("Student Details!")

	var name string = "Aditya Mayukh Som"
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Length of name is %d\n", utf8.RuneCountInString(name))

	var age int = 20
	fmt.Printf("My age is %d\n", age)

	var gender rune = 'M'
	if gender == 'M' {
		fmt.Println("Male")
	} else {
		fmt.Println("Female")
	}

	var n int = 10000
	var testSliceWithoutCapacity []int = []int{}
	var testSliceWithCapacity []int = make([]int, 0, n)

	var durationWithoutCapacity time.Duration = measureSliceAppendTime(testSliceWithoutCapacity, n)
	var durationWithCapacity time.Duration = measureSliceAppendTime(testSliceWithCapacity, n)

	fmt.Printf("Time without preallocation is %v\n", durationWithoutCapacity)
	fmt.Printf("Time with preallocation is %v\n", durationWithCapacity)

	var speedup float64 = float64(durationWithoutCapacity.Nanoseconds()) / float64(durationWithCapacity.Nanoseconds())
	fmt.Printf("Speedup is %f\n", speedup)
}

func measureSliceAppendTime(slice []int, n int) time.Duration {
	var initTime time.Time = time.Now()

	for i := 0; i < n; i++ {
		slice = append(slice, 69)
	}

	var duration time.Duration = time.Since(initTime)
	return duration

}
