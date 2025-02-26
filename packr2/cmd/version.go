package cmd

import (
	"fmt"

	packr "github.com/arconomy/packr"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows packr version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(packr.Version)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
