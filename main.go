package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func parseFile() {
	buffer := make([]byte, 8)
	bufferPointer := 0
	fileName := "messages.txt"

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
		fmt.Printf("read: %s\n", buffer[:read])
		if err == io.EOF {
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

    Remove the printing of "I hope I get the job!".

Instead, your program will now read messages.txt 8 bytes at a time and print that data back to stdout in 8 byte chunks.
Here's some pseudocode:

    os.Open messages.txt for reading.
    While there is still data in the file:
        Read 8 bytes from the file into a slice of bytes.
        Print the 8 bytes as text to stdout in this format: read: %s\n
        When you finally get an io.EOF error, exit the program.

*/
