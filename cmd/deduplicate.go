package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gkwa/newtoniansheep/core"
	"github.com/spf13/cobra"
)

var deduplicateCmd = &cobra.Command{
	Use:   "deduplicate [file]",
	Short: "Deduplicate image links in a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		fileHandler := core.NewFileHandler()
		processor := core.NewProcessor()
		deduplicator := core.NewDeduplicator(logger, fileHandler, processor)

		inputPath := args[0]
		absInputPath, err := filepath.Abs(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get absolute input file path: %v\n", err)
			absInputPath = inputPath
		}

		duplicatesRemoved, err := deduplicator.ProcessFile(absInputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to process file: %v\n", err)
			os.Exit(1)
		}

		metadata, err := core.GetFileMetadata(absInputPath, duplicatesRemoved)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get file metadata: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(metadata.String())
	},
}

func init() {
	rootCmd.AddCommand(deduplicateCmd)
}
