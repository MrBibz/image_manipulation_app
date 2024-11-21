package main

import (
	"fmt"
	"time"
)

// Hash function that takes a string input and returns a 64-bit hash value
func Hash(input string) uint64 {
	var hash uint64 = 0xcbf29ce484222325 // Initial hash value (FNV-1a offset basis)
	var prime uint64 = 0x100000001b3     // FNV-1a prime

	// Iterate over each byte in the input string
	for i := 0; i < len(input); i++ {
		hash ^= uint64(input[i])           // XOR the hash with the byte value
		hash *= prime                      // Multiply the hash by the FNV-1a prime
		hash = (hash << 13) | (hash >> 51) // Rotate the hash left by 13 bits
		hash ^= hash >> 7                  // XOR the hash with itself shifted right by 7 bits
	}

	return hash // Return the final hash value
}

func main() {
	data := "Here is a string of characters in order to test Go's execution speed." // Input data to be hashed

	var hashValue uint64 // Variable to store the hash value
	start := time.Now()  // Record the start time

	// Loop to hash the data 100,000 times
	for i := 0; i < 100000; i++ {
		hashValue = Hash(data) // Compute the hash value
	}
	elapsed := time.Since(start) // Calculate the elapsed time

	// Print the final hash value and the execution time
	fmt.Println("Hash value:", hashValue)
	fmt.Println("Execution time:", elapsed)
}
