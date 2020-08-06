package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Need a CSV file as an argument, e.g., csv2text somefile.csv")
	}
	infile, err := os.Open("sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	r := csv.NewReader(infile)
	// Read first line to get column names
	columnNames, err := r.Read()
	if err != nil {
		log.Fatalf("failed reading first line for column names: %v", err)
	}

	lineNumber := 2 // we already read the first line
	for {
		l, err := r.Read()
		if err == io.EOF {
			fmt.Println("finished reading file")
			break
		}
		if len(l) != len(columnNames) {
			log.Fatalf("line %d contained %d fields which does not match the %d field titles", lineNumber, len(l), len(columnNames))
		}
		err = writeFile(columnNames, l)
		if err != nil {
			log.Fatalf("failed to write file for line %d: %v", lineNumber, err)
		}
		lineNumber++
	}
}

func writeFile(cols, vals []string) error {
	out, err := os.Create(vals[0]+".txt")
	if err != nil {
		return fmt.Errorf("unable to open file '%s' for writing: %v", vals[0]+".txt", err)
	}
	defer out.Close()

	for i := 0; i < len(cols); i++ {
		fmt.Fprintf(out, "%s: %s\n", cols[i], vals[i])
	}
	return nil
}