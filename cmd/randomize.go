package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/gkwa/newtoniansheep/core"
)

var randomizeCmd = &cobra.Command{
	Use:   "randomize [file]",
	Short: "Randomize image links in a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		fileHandler := core.NewFileHandler()
		processor := core.NewRandomizer()
		randomizer := core.NewLinkRandomizer(logger, fileHandler, processor)

		inputPath := args[0]
		absInputPath, err := filepath.Abs(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get absolute input file path: %v\n", err)
			absInputPath = inputPath
		}

		outputPath := core.GetRandomizedFilePath(absInputPath)
		absOutputPath, err := filepath.Abs(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get absolute output file path: %v\n", err)
			absOutputPath = outputPath
		}

		err = randomizer.ProcessFile(absInputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to process file: %v\n", err)
			os.Exit(1)
		}

		metadata, err := core.GetFileMetadata(absOutputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get file metadata: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(metadata.String())
	},
}

func init() {
	rootCmd.AddCommand(randomizeCmd)
}
