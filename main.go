package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

var (
	PrintErrorF = color.New(color.FgRed).FprintfFunc()
	PrintError  = color.New(color.FgRed).PrintFunc()
)

func main() {

	defer recoverFn()
	outputFilePath := flag.String("output-file", "StdOut", `Where the data will be wrote. \nBy default this will be displayed on Stdout to send data througth >> bash statement`)
	separator := flag.String("separ", ",", "What will separate all data")
	dataFilePath := flag.String("file-path", "", "Data file path to append")

	flag.Parse()

	if _, e := os.Stat(*dataFilePath); os.IsNotExist(e) {
		PrintErrorF(os.Stdout, "file-path '%s' does not exists", *dataFilePath)
		return
	}

	dataFile, e := os.Open(*dataFilePath)
	if e != nil {
		PrintError("Error opening file-path")
		return
	}
	defer dataFile.Close()

	outputWriter, close := getOutputWriter(*outputFilePath)
	if close != nil {
		defer close()
	}

	sc := bufio.NewScanner(dataFile)
	for sc.Scan() {
		dataToWrite := fmt.Sprintf("%s%s", sc.Text(), *separator)
		fmt.Fprint(outputWriter, dataToWrite)
	}

}

func getOutputWriter(outputFilePath string) (io.Writer, func() error) {
	if outputFilePath == "StdOut" {
		return os.Stdout, nil
	}

	outputWriter, e := os.Create(outputFilePath)
	if e != nil {
		panic("Error opening output-file")
	}

	return outputWriter, outputWriter.Close

}

func recoverFn() {
	message := recover()
	if message != nil {
		color.Red("An unexpected error has occured :/")
		return
	}
	color.Green("Ready!")
}
