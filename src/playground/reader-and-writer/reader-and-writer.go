package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// myReader()
	// readFromPkg()
	// readFromTerminal()
	writerFromPkg()
}

func myReader() {
	// string to read
	reader := strings.NewReader("Rafael")
	// new String builder
	var newString strings.Builder
	// empty buffer
	buffer := make([]byte, 4)

	for {
		// get how many bytes has been read
		numBytes, err := reader.Read(buffer)
		// place bytes into chunk
		chunk := buffer[:numBytes]
		// write chunk into new string
		newString.Write(chunk)
		// print it
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		// stop loop when string ends
		if err == io.EOF {
			break
		}
	}
	// prints new string
	fmt.Printf("%v\n", newString.String())
}

func readFromPkg() {
	// String to read
	source := strings.NewReader("Rafael")
	// buffered reader, supplying the string to read
	buffered := bufio.NewReader(source)
	// Split on the end of the line
	newString, err := buffered.ReadString('\n')
	// stops reading when file ends
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("Something went wrong")
	}
}

func readFromTerminal() {
	// to run:
	// 	printf "these\nare\nsome\n0\nwords" \
	// | go run playground/reader-and-writer/reader-and-writer.go

	// Read lines from stdin
	scanner := bufio.NewScanner(os.Stdin)
	// create buffer
	lines := make([]string, 0, 5)
	// show inputs
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}
	for _, line := range lines {
		fmt.Printf("Line: %v\n", line)
	}
	if scanner.Text() == "0" {
		os.Exit(0)
	}
}

func writerFromPkg() {
	buffer := bytes.NewBufferString("")
	numBytes, err := buffer.WriteString("Rafael")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wrote %v bytes: %v", numBytes, buffer)
	}
}
