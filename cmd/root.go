package cmd

import (
	"fmt"
	"os"

	"github.com/enrico-kaack/github-link-formater/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-link-formatter",
	Short: "A CLI tool to format GitHub issue and PR links",
	Long:  `This tool takes a GitHub issue or pull request URL and formats it into a markdown link with relevant details.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a GitHub issue or pull request URL.")
			return
		}
		url := args[0]

		formatted, err := internal.FormatGHLink(url)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println(formatted)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
