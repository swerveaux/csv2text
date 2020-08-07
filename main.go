package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var inputFilename string
	var outputDirectory string

	flag.StringVar(&inputFilename, "in", "", "path to input csv file: e.g., 'file.csv' or '/Users/someone/csvfiles/file.csv")
	flag.StringVar(&outputDirectory, "outdir", ".", "path to directory to create output files, defaults to the current working directory ('.')")
	flag.Parse()

	if inputFilename == "" {
		log.Fatalf("You must specify an input filename with --in")
	}

	infile, err := os.Open(inputFilename)
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
		err = writeFile(columnNames, l, outputDirectory)
		if err != nil {
			log.Fatalf("failed to write file for line %d: %v", lineNumber, err)
		}
		lineNumber++
	}
}

func writeFile(cols, vals []string, outputFileDir string) error {
	out, err := os.Create(filepath.Join(outputFileDir, vals[0]+" "+vals[1]+".txt"))
	if err != nil {
		return fmt.Errorf("unable to open file '%s' for writing: %v", vals[0]+".txt", err)
	}
	defer out.Close()

	for i := 0; i < len(cols); i++ {
		fmt.Fprintf(out, "%s: %s\n", cols[i], vals[i])
	}
	return nil
}
