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
	"unicode"
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

func numberOfWords(fileName string) int {

	fileData, err := os.ReadFile(fileName)
	var numberOfWords int

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, v := range string(fileData) {
		if unicode.IsSpace(v) {
			numberOfWords++
		}

	}

	return numberOfWords
}

func main() {
	byteCmd := flag.Bool("b", false, "print file byte number")
	lineCmd := flag.Bool("l", false, "return number of lines in a file")
	numberCmd := flag.Bool("w", false, "return number of words in a file")
	// charCmd := flag.Bool("m", false, "return number of chars in a file")

	flag.Parse()
	fileName := flag.Args() //Checks for the last argument and assigns that as the file name

	// Make sure there is atleast one (the file argument) passed
	if len(os.Args[1:]) < 1 {
		fmt.Println("Please pass an argument")
		os.Exit(1)
	}

	if len(fileName) < 1 {
		fmt.Println("Please enter a file")
		os.Exit(1)
	}

	if *byteCmd {
		fmt.Println(numberOfBytes(fileName[0]))
	}

	if *lineCmd {
		fmt.Println(numberOfLines(fileName[0]))
	}

	if *numberCmd {
		fmt.Println(numberOfWords(fileName[0]))
	}
}
