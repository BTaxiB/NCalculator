/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
