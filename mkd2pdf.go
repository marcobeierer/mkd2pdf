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
	count, err := fmt.Scanf("%d", &filenumber)
	if err != nil {
		panic(err)
	}

	if count != 1 || filenumber > len(filenames) {

		fmt.Println("Invalid input. Please try again.")
		return
	}

	selectedFilename := filenames[filenumber]
	pdfFilename := strings.Replace(selectedFilename, ".mkd", ".pdf", -1) // TODO use regexp to make     sure that at the end

	_, err = os.Stat(pdfFilename)
	if err == nil {

		fmt.Printf("File %s already exists.\nDo you want to overwrite the file? [yes/no]\n", pdfFilename)

		var overwrite string
		_, err := fmt.Scanf("%s", &overwrite)
		if err != nil {
			panic(err)
		}

		if overwrite != "yes" {
			return
		}
	}

	cmd := exec.Command("pandoc", selectedFilename, "-o", pdfFilename)

	fmt.Print("Do you need a table of contents? [yes/no]\n")

	var toc string
	_, err = fmt.Scanf("%s", &toc)
	if err != nil {
		panic(err)
	}

	if toc == "yes" {
		cmd.Args = append(cmd.Args, "--toc")
	}

	fmt.Printf("Conversation started. Please wait.\n")

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s successfully converted to %s.\n", selectedFilename, pdfFilename)
	return
}
