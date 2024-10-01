package cmd

import (
	"fmt"
	"os"

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
		manager := core.NewDeduplicateManager(logger, fileHandler, processor)

		result, err := manager.Deduplicate(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(deduplicateCmd)
}
