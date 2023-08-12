package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Define a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Define a boolean flag -l to count bytes instead of words
	bytes := flag.Bool("b", false, "Count bytes")
	// Parse the flags provided by the user
	flag.Parse()
	// Calling the count function to count the number of words received from the Standard Input
	// and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set we want to count the words so we define
	// the scanner split type to words (default is split by line)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Define a counter
	wc := 0

	// For ever word counted, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
