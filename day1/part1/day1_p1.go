package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../day1_input.txt")
	r := bufio.NewReader(f)

	var total int64
	for {
		// Read in the next line in the file,
		// if we are at the end break out
		cmdLine, err := r.ReadBytes('\n')

		// Setup the command and value for computation
		cmdLineLen := len(cmdLine)
		command := cmdLine[0]
		// The integer we want to compute will be between this range
		stringValue := string(cmdLine[1:(cmdLineLen - 2)])
		if err != nil && err.Error() == "EOF" {
			stringValue = string(cmdLine[1:])
		}
		value, _ := strconv.ParseInt(stringValue, 0, 64)

		switch command {
		case 43:
			total += value
		case 45:
			total -= value
		}

		if err != nil && err.Error() == "EOF" {
			break
		}
	}

	f.Close()
	fmt.Println(total)
}
