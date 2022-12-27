package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	PrintErrorF    = color.New(color.FgRed).FprintfFunc()
	PrintError     = color.New(color.FgRed).PrintFunc()
	outputFilePath string
	separator      string
	dataFilePath   string
)

var RootCmd = &cobra.Command{
	Use: "join-matic [flags] [args...]",
	Example: `
join-matic  -f "/path/to/file/data.csv" -o /path/to/output.txt -s /
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if _, e := os.Stat(dataFilePath); os.IsNotExist(e) {
			PrintErrorF(os.Stdout, "file-path '%s' does not exists\n", dataFilePath)
			return
		}

		dataFile, e := os.Open(dataFilePath)
		if e != nil {
			PrintError("Error opening file-path")
			return
		}
		defer dataFile.Close()

		outputWriter, close, err := getOutputWriter(outputFilePath)
		if err != nil {
			PrintError(err)
		}
		if close != nil {
			defer close()
		}

		sc := bufio.NewScanner(dataFile)
		for sc.Scan() {
			dataToWrite := fmt.Sprintf("%s%s", sc.Text(), separator)
			fmt.Fprint(outputWriter, dataToWrite)
		}

		color.Green("\nReady!")
	},
}

func getOutputWriter(outputFilePath string) (io.Writer, func() error, error) {
	if outputFilePath == "StdOut" {
		return os.Stdout, nil, nil
	}

	outputWriter, e := os.Create(outputFilePath)
	if e != nil {
		return nil, nil, fmt.Errorf("error opening output-file. Details: %w", e)
	}

	return outputWriter, outputWriter.Close, nil

}

func init() {
	RootCmd.Flags().StringVarP(&outputFilePath, "output-file", "o", "StdOut", "Where the data will be wrote. By default this will be displayed on Stdout to send data througth >> bash statement")
	RootCmd.Flags().StringVarP(&separator, "separ", "s", ",", "What will separate all data")
	RootCmd.Flags().StringVarP(&dataFilePath, "file-path", "f", "", "Data file path to append")
}
