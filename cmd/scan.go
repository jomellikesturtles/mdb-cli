package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var list1 = [2]string{"/Users/jommel/Downloads/Videos", "/Users/jommel/Downloads/torrents"}
var sampleFolder = "~/Downloads/torrents"

var scanCmd = &cobra.Command{

	Use:   "scan [path]",
	Short: "Scan a directory for movie files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		root := args[0]

		fmt.Println(len(list1))
		for i := 0; i < len(list1); i++ {
			// fmt.Println()
			fmt.Printf("Scanning directory: %s\n", list1[i])
		}
		return

		validExtensions := map[string]bool{
			".mp4":  true,
			".mkv":  true,
			".mpeg": true,
			".avi":  true,
			".wmv":  true,
			".mpg":  true}

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

func readDirectory(startPath string) {
	// if exists
	// fmt.Println("Directory does not exist: %s", startPath)
	// return
	filesList := []string{}
	for i := 0; i < len(filesList); i++ {
		filename :=
			readDirectory(filename)
	}
}
func isVideoFile(filename string) {

}
func addToList(folderPath string, fileName string) {
	fmt.Println("Adding to List: %s", fileName)
	saveToDb("")
}

func saveToDb(obj string) {
	fmt.Println("Saving to DB: %s", obj)

}

func init() {
	rootCmd.AddCommand(scanCmd)
}
