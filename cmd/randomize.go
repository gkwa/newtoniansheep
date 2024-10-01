package cmd

import (
	"fmt"
	"os"

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
		manager := core.NewRandomizeManager(logger, fileHandler, processor)

		result, err := manager.Randomize(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(randomizeCmd)
}
