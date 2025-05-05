package main

import (
	"fmt"
)

func main() {
	// Your code goes here

	m := make(map[string]int)

	// Insert or update
	m["apple"] = 5    // set key "apple" to 5
	m["apple"] = 7    // update to 7
	
	// Retrieve (zero‑value if missing)
	count := m["banana"] // 0, but banana may not exist!
	
	// Delete
	// delete(m, "apple")   // removes key "apple"; safe if not present
	
	// Length

	fmt.Println(len(m), count)  // number of keys currently in map
	
	k , ok := m["apple"] // ok is false if key "apple" does not exist
	fmt.Println(k, ok) // 0 false
}
