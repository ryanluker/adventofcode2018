package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var freq int64
	prevTotals := make(map[int64]int)

	for {
		f, _ := os.Open("../day1_input.txt")
		r := bufio.NewReader(f)
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
				freq += value
			case 45:
				freq -= value
			}

			// see if the freq has already been seen
			_, prs := prevTotals[freq]
			if prs == true {
				f.Close()
				fmt.Println("Double found! ", freq)
				panic("stop")
			}

			prevTotals[freq] = 1

			if err != nil && err.Error() == "EOF" {
				fmt.Println("No doubles found! len: ", len(prevTotals))
				f.Close()
				break
			}
		}
	}
}
