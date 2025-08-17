package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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
		switch {
		case attribute == "gopher" && name == "sakanakun":
			ss := []string{"Go", "ぎょ"}
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(ss), func(i, j int) { ss[i], ss[j] = ss[j], ss[i] })

			cnt := strings.Count(message, "ご")
			if cnt > 0 {
				for i := 0; i < cnt; i++ {
					message = strings.Replace(message, "ご", ss[i%len(ss)], 1)
				}
			}
			fmt.Println(message)
		case attribute == "gopher":
			fmt.Println(strings.ReplaceAll(message, "ご", "Go"))
		case name == "sakanakun":
			fmt.Println(strings.ReplaceAll(message, "ご", "ギョ"))
		default:
			fmt.Println(message)
		}
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
