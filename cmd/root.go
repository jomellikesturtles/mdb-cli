package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "mdb",
	Short: "MDB CLI - Manage your local media library",
	Long:  `A high-performance CLI companion for the MDB Platform to scan, identify, and sync local media.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mdb.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		log.Println("cfgFile: ", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		home, _ := os.UserHomeDir()
		log.Println("Home: ", home)
		viper.AddConfigPath(home)
		viper.SetConfigName(".mdb")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
