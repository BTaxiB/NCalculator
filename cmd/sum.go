package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// sumCmd represents the sum command
var sumCmd = &cobra.Command{
	Use:   "sum",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var result float64

		result, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		for i, number := range args {

			if i == 0 {
				continue
			}
			temp, err := strconv.ParseFloat(number, 64)
			if err != nil {
				log.Fatal(err)
			}

			result *= temp
		}

		log.Printf("result: %f", result)
	},
}

func init() {
	rootCmd.AddCommand(sumCmd)
}
