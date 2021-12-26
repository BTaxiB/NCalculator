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
