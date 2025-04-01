package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parseFile() {
	buffer := make([]byte, 8)
	bufferPointer := 0
	fileName := "messages.txt"
	currentLine := ""

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	if err != nil {
		log.Fatal(err)
	}

	for {
		read, err := file.ReadAt(buffer, int64(bufferPointer))

		currentChunk := string(buffer[:read])
		currentLine += currentChunk

		if strings.Contains(currentChunk, "\n") {
			splitLine := strings.Split(currentLine, "\n")
			for i := 0; i < len(splitLine)-1; i++ {
				fmt.Printf("read: %s\n", splitLine[i])
			}
			currentLine = splitLine[len(splitLine)-1]
		}

		if err == io.EOF {
			if currentLine != "" {
				fmt.Printf("read: %s\n", currentLine)
			}
			break
		}
		bufferPointer += read

	}
}

func main() {
	parseFile()
}

/*

Assignment

Let's update our code to continue to read 8 bytes at a time, but now let's format the output line by line. It should now
print the data in the same format:

read: %s\n

But now %s represents an entire line of the input file. We're decoupling how we present and parse (full lines) the data
from how we read it (8 bytes at a time).

Here's what I did:

    Create a string variable to hold the contents of the "current line" of the file. It needs to persist between reads
(loop iterations).
    After reading 8 bytes, split the data on newlines (\n) to create a slice of strings - let's call these split
sections "parts". There will typically only be one or two "parts" because we're only reading 8 bytes at a time.
    For each part except the last one, print a line to the console in this format:

read: LINE

Where LINE is the "current line" we've aggregated so far plus the current "part". Then reset the "current line" variable
to an empty string. Note that if we only have one "part", we don't need to print, as we have not reached a new line yet.

    Add the last "part" to the "current line" variable. Repeat until you reach the end of the file.
    Once you're done reading the file, if there's anything left in the "current line" variable, print it in the same
read: LINE format.


*/
