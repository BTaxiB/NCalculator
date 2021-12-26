package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Calculator interface {
	AddNumbers(numbers ...float64) float64
	SubstractNumbers(numbers ...float64) float64
	DivideNumbers(dividend float64, divisor float64) (float64, error)
	MultiplyNumbers(numbers ...float64) float64
}

type NGCalculator struct {
	Result float64
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ngcalculator",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	calculator := NewCalculator()
}

func NewCalculator() Calculator {
	return NGCalculator{}
}

func (a NGCalculator) AddNumbers(numbers ...float64) float64 {
	var result float64

	for _, number := range numbers {
		result += number
	}

	return result
}

func (a NGCalculator) SubstractNumbers(numbers ...float64) float64 {
	var result float64

	for _, number := range numbers {
		result -= number
	}

	return result
}

func (a NGCalculator) DivideNumbers(dividend float64, divisor float64) (float64, error) {
	var result float64

	result = dividend / divisor

	if result < 0 {
		return 0, fmt.Errorf("failed to divide %f with %f", dividend, divisor)
	}

	return result, nil
}

func (a NGCalculator) MultiplyNumbers(numbers ...float64) float64 {
	var result float64

	for _, number := range numbers {
		result *= number
	}

	return result
}
