package cmd

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get a resource",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
