package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var result float64

		for _, number := range args {
			temp, err := strconv.ParseFloat(number, 64)
			if err != nil {
				log.Fatal(err)
			}

			result += temp
		}

		log.Printf("result: %f", result)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
