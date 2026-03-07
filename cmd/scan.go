package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan a directory for movie files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := args[0]
		// root = "/Users/jommel/Downloads/torrents"
		fmt.Printf("Scanning directory: %s...\n", root)

		extensions := map[string]bool{".mp4": true, ".mkv": true, ".avi": true, ".mpeg4": true}

		foundList := []string{}
		err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() && extensions[filepath.Ext(path)] {
				// In a real app, you'd call internal/metadata here
				foundList = append(foundList, d.Name())
				fmt.Printf("Found: %s\n", d.Name())
			}
			return nil
		})
		fmt.Printf("Found List: %s\n", strings.Join(foundList, ","))

		if err != nil {
			fmt.Printf("Error scanning: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
