//  Step 1  - (flag -c) return number of bytes in a file
// step 2  - (flag -l) return number of lines in a file
// step 3 - (flag -w) return number of words in a file
// step 4 - (flag -m) return number of characters in a file
// step 5 - (flag none) return all of the above
// NB print the file name being parsed

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func numberOfBytes(fileName string) int64 {
	openedFile, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)

	}
	return openedFile.Size()

}

func numberOfLines(fileName string) int {
	fileData, err := os.ReadFile(fileName)
	var numberOfLines int = 0
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileData), "\n")

	for numberOfLines = range lines {
	}

	return numberOfLines
}

// TOO:FIX = Not accurate to the wc linux util
func numberOfWords(fileName string) int {

	fileData, err := os.ReadFile(fileName)
	var numberOfWords int

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for i := range fileData {
		if fileData[i] == 32 || fileData[i] == 10 || fileData[i] == 42 { //Check the ascii char
			continue
		}
		numberOfWords++
	}
	return numberOfWords
}

func numberOfChars(fileName string) int {
	fileData, err := os.ReadFile(fileName)
	var numberOfChars int

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for range string(fileData) {
		numberOfChars++ // chars inclusive of end of line and white
	}

	return numberOfChars
}

func runAllFlags(fileName string) (int64, int, int, int) {
	fmt.Printf("--bytes=%d --lines=%d  --words=%d --chars=%d \n", numberOfBytes(fileName), numberOfLines(fileName), numberOfWords(fileName), numberOfChars(fileName))
	return numberOfBytes(fileName), numberOfLines(fileName), numberOfWords(fileName), numberOfChars(fileName)

}

func main() {
	byteCmd := flag.Bool("b", false, "print file byte number")
	lineCmd := flag.Bool("l", false, "return number of lines in a file")
	numberCmd := flag.Bool("w", false, "return number of words in a file")
	charCmd := flag.Bool("m", false, "return number of chars in a file")

	flag.Parse()
	fileName := flag.Args() //Checks for the last argument and assigns that as the file name

	// Make sure there is atleast one (the file argument) passed
	if len(fileName) < 1 {
		fmt.Println("Please enter a file")
		os.Exit(1)
	}

	if len(os.Args[1:]) >= 2 {
		if *byteCmd {
			fmt.Printf("--bytes=%d\n", numberOfBytes(fileName[0]))
		}

		if *lineCmd {
			fmt.Printf("--lines=%d\n", numberOfLines(fileName[0]))
		}

		if *numberCmd {
			fmt.Printf("--words=%d\n", numberOfWords(fileName[0]))
		}

		if *charCmd {
			fmt.Printf("--chars=%d\n", numberOfChars(fileName[0]))
		}

	} else {
		runAllFlags(fileName[0])
	}

}
