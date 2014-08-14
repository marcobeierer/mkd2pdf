package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	fmt.Println("Available markdown files:")

	filenames, err := filepath.Glob("*.mkd")
	if err != nil {
		panic(err)
	}

	if len(filenames) < 1 {
		fmt.Println("No files available.")
		return
	}

	for index, filename := range filenames {
		fmt.Printf("[%d] %s\n", index, filename)
	}

	fmt.Print("\nSelect a file to convert: ")

	var filenumber int
	count, err := fmt.Scanf("%d\n", &filenumber)
	if err != nil {
		panic(err)
	}

	if count != 1 {
		panic("invalid input")
	}

	selectedFilename := filenames[filenumber]
	pdfFilename := strings.Replace(selectedFilename, ".mkd", ".pdf", -1) // TODO use regexp to make sure that at the end

	_, err = os.Stat(pdfFilename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", pdfFilename)
		return
	}

	cmd := exec.Command("pandoc", "--toc", selectedFilename, "-o", pdfFilename)

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s successfully converted to %s.\n", selectedFilename, pdfFilename)
}
