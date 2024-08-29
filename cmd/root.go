package cmd

import "os"
import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "forge",
	Short: "A project generator built in Golang!",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
