// entry point and global flags
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns current version",
	Run: func(cmd *cobra.Command, args []string) {
		// viper.WriteConfig()

		// log.Println(viper.Get("auth.token"))
		fmt.Println("1.0.0")
		// fmt.Println(viper.ReadInConfig())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
