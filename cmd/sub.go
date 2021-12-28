package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
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

			result -= temp
		}

		log.Printf("result: %f", result)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
