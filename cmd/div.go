package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var result float64

		dividend, err := strconv.ParseFloat(args[0], 64)
		divisor, err := strconv.ParseFloat(args[1], 64)

		if err != nil {
			log.Fatal(err)
		}

		result = dividend / divisor

		if result < 0 {
			log.Fatalf("failed to divide %f with %f result: %f", dividend, divisor, result)
		}

		log.Printf("result: %f", result)
	},
}

func init() {
	rootCmd.AddCommand(divCmd)

}
