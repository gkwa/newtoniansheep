package cmd

import (
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

		err := deduplicator.ProcessFile(args[0])
		if err != nil {
			logger.Error(err, "Failed to process file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(deduplicateCmd)
}
