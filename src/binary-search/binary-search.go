package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr *[1000]int, search int) int {
	var minIndex int = 0         // Minimum index
	var maxIndex int = len(*arr) // Maximum index
	for minIndex != maxIndex {   // Min and max index will be squeezed together until they become the same
		var index int = (minIndex + maxIndex) / 2 // Get the middle of the min and max index
		var value int = (*arr)[index]             // Value at that index
		if value == search {                      // If value matches search return index
			return index
		} else if value < search { // If value is less than search, min index becomes index + 1
			minIndex = index + 1
		} else { // If value is more than search max index becomes index
			maxIndex = index
		}
	}
	return minIndex // Min and max are same, so return either
}

func main() {
	rand.Seed(time.Now().UnixMilli()) // Random seed set to unix time in milliseconds
	var data [1000]int                // Create variable "data" that is an integer array that holds 1,000 elements
	for i := 0; i < len(data); i++ {  // Fill data with random numbers from 0 to 1,000,000
		data[i] = rand.Intn(1_000_000)
	}
	sort.Ints(data[:])                                 // Sort array, because sort.Ints takes a slice, [:] is used to convert array to slice
	var randomElement int = data[rand.Intn(len(data))] // Get random element from array to search for
	fmt.Println(binarySearch(&data, randomElement))    // Run binary search function with array and search value, and print it
}
