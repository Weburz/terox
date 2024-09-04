package cmd

import (
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a project using a template.",
	Example: `forge create "Weburz/simple-website"`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Matches the pattern - "Weburz/repoforge" and so on
		pattern := "^[a-zA-Z0-9_-]+/[a-zA-Z0-9_-]+$"

		// Compiles the pattern into the binary before a string match can be
		// performed
		re, err := regexp.Compile(pattern)

		// Throw and error and exit the execution if the pattern is invalid
		if err != nil {
			rootCmd.PrintErrf("Invalid regex pattern: %s\n", err)
			os.Exit(1)
		}

		// Match the accepted regex pattern with the argument passed to the
		// command
		match := re.MatchString(args[0])

		// Throw error and exit the execution if the URL passed to the command
		// is invalid
		if !match {
			rootCmd.PrintErrf("Error: %s is an invalid template URL\n", args[0])
			os.Exit(1)
		}

		// TODO: Perform the rest of the logic if there were no errors
		rootCmd.Printf("%s matches the pattern %s\n", args[0], pattern)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
