package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gkwa/newtoniansheep/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of newtoniansheep",
	Long:  `All software has versions. This is newtoniansheep's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
