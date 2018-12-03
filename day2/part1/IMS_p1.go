package main

import (
	"bufio"
	"fmt"
	"os"
)

// CreateUniqueMap creates a unique mapping out of a line of bytes
func CreateUniqueMap(currLine []byte) map[string]int {
	currMap := make(map[string]int)
	for {
		var currChar byte
		currChar, currLine = currLine[0], currLine[1:]

		// Check new line to know when to stop
		if currChar == 10 {
			break
		}
		currMap[string(currChar)]++
	}
	return currMap
}

// FindCopies looks through the map for the first value that matches howMany
// and returns 1 or 0 if none are found
func FindCopies(currMap map[string]int, howMany int) int {
	for _, value := range currMap {
		if value == howMany {
			return 1
		}
	}
	return 0
}

// Inventory management system
// looks at each inventory code work (one per line) looking for double and
// or triple continuous exact characters
func main() {
	f, _ := os.Open("./day2_input.txt")
	r := bufio.NewReader(f)

	var triples int
	var doubles int
	for {
		currLine, err := r.ReadBytes('\n')
		if err != nil {
			// Add new line to prevent panic
			currLine = append(currLine, '\n')
		}

		currMap := CreateUniqueMap(currLine)
		doubles += FindCopies(currMap, 2)
		triples += FindCopies(currMap, 3)

		if err != nil {
			break
		}
	}

	hash := doubles * triples
	f.Close()
	fmt.Printf("hash for the current inventory store list: %d\n", hash)
}
