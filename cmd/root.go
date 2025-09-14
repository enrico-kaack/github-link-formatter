package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/enrico-kaack/github-link-formatter/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-link-formatter <github issue or pr url>",
	Short: "A CLI tool to format GitHub issue and PR links",
	Long:  `This tool formats GitHub links for issues and PRs to repos into readable formats (like Markdown links) by querying the GitHub API for more information.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a GitHub issue or pull request URL as an argument")
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
