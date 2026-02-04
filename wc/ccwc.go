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
)

func numberOfBytes(fileName string) int64 {
	openedFile, err := os.Stat(fileName)

	if err != nil {
		log.Fatal(err)

	}
	return openedFile.Size()

}

func main() {
	byteCmd := flag.Bool("b", false, "print file byte number")
	// lineCmd := flag.Bool("l", false, "return number of lines in a file")
	// numberCmd := flag.Bool("w", false, "return number of words in a file")
	// charCmd := flag.Bool("m", false, "return number of chars in a file")

	flag.Parse()
	fileNames := flag.Args()

	// if len(os.Args[1:]) < 1 {
	// 	fmt.Println("no flag entered")
	// 	os.Exit(1)
	// }

	if *byteCmd {
		fmt.Println(numberOfBytes(fileNames[0]))
	}
}
