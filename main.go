package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	attribute string
	name      string
	message   string
)

var rootCmd = &cobra.Command{
	Use:   "sakanapher",
	Short: "A CLI tool to convert messages like a fish-loving gopher.",
	Long:  `sakanapher is a CLI application that replaces specific characters in a message based on attributes. `,
	Run: func(cmd *cobra.Command, args []string) {
		s := Sakanize(message, attribute, name)
		fmt.Println(s)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&attribute, "attribute", "a", "", "your attribute [gopher|other]")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "your name")
	rootCmd.Flags().StringVarP(&message, "message", "m", "おはようございます", "message of wish to convert")
}

func main() {
	Execute()
}
