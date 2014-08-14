package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	selectedFilename := selectFileToConvert()
	pdfFilename := strings.Replace(selectedFilename, ".mkd", ".pdf", -1) // TODO use regexp to make     sure that at the end

	overwriteExistingFile(pdfFilename)

	cmd := exec.Command("pandoc", selectedFilename, "-o", pdfFilename)

	tableOfContents(cmd)

	fmt.Printf("\nConversation started. Please wait.\n")

	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s successfully converted to %s.\n", selectedFilename, pdfFilename)
	return
}

func selectFileToConvert() string {

	fmt.Println("Available markdown files:")

	filenames, err := filepath.Glob("*.mkd")
	if err != nil {
		panic(err)
	}

	if len(filenames) < 1 {

		fmt.Println("No files available.")
		os.Exit(1)
	}

	for index, filename := range filenames {
		fmt.Printf("[%d] %s\n", index, filename)
	}

	fmt.Print("\nSelect a file to convert: ")

	var filenumber int
	count, err := fmt.Scanf("%d", &filenumber)
	if err != nil || count != 1 || filenumber > len(filenames) {

		fmt.Println("Invalid input. Please try again.")
		os.Exit(1)
	}

	selectedFilename := filenames[filenumber]

	return selectedFilename
}

func overwriteExistingFile(pdfFilename string) {

	_, err := os.Stat(pdfFilename)
	if err == nil {

		fmt.Printf("\nFile %s already exists. Do you want to overwrite the file? [yes/no]\n", pdfFilename)

		var overwrite string
		_, err := fmt.Scanf("%s", &overwrite)
		if err != nil {
			panic(err)
		}

		if overwrite != "yes" {
			os.Exit(0)
		}
	}

	return
}

func tableOfContents(command *exec.Cmd) {

	fmt.Print("\nDo you need a table of contents? [yes/no]\n")

	var toc string
	_, err := fmt.Scanf("%s", &toc)
	if err != nil {
		panic(err)
	}

	if toc == "yes" {
		command.Args = append(command.Args, "--table-of-contents")
	}

	return
}
