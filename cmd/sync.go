package cmd

import "github.com/spf13/cobra"

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync metadata and user data",
	Args:  cobra.MinimumNArgs(0),
	// Run: func(cmd)
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
