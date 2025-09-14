package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/enrico-kaack/github-link-formater/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-link-formatter",
	Short: "A CLI tool to format GitHub issue and PR links",
	Long:  `This tool takes a GitHub issue or pull request URL and formats it into a markdown link with relevant details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			fmt.Println("Please provide a GitHub issue or pull request URL.")
			return nil
		}
		url := args[0]

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir := filepath.Join(home, ".github-link-formatter")

		formatted, err := internal.FormatGHLink(url, configDir)
		if err != nil {
			return err
		}

		fmt.Println(formatted)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
