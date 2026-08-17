package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var scanCmd3 = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for movie files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := args[0]
		fmt.Printf("Scanning directory: %s...\n", root)

		extensions := map[string]bool{".mp4": true, ".mkv": true, ".avi": true}

		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && extensions[filepath.Ext(path)] {
				// In a real app, you'd call internal/metadata here
				fmt.Printf("Found: %s\n", d.Name())
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error scanning: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
