package cmd

import (
	"fmt"

	"github.com/jon4hz/ethconvert/pkg/ethconvert"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var rootCmdFlags struct {
	input  string
	output string
	amount string
}

var rootCmd = &cobra.Command{
	Use:   "ethconvert",
	Short: "convert ethereum units",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 1:
			rootCmdFlags.amount = args[0]
		case 2:
			rootCmdFlags.amount = args[0]
			rootCmdFlags.input = args[1]
		case 3:
			rootCmdFlags.amount = args[0]
			rootCmdFlags.input = args[1]
			rootCmdFlags.output = args[2]
		}
		amount, err := decimal.NewFromString(rootCmdFlags.amount)
		if err != nil {
			return err
		}
		o, err := ethconvert.Convert(amount, rootCmdFlags.input, rootCmdFlags.output)
		if err != nil {
			return err
		}
		fmt.Printf("%s %s\n", o, rootCmdFlags.output)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVarP(&rootCmdFlags.input, "input", "i", "ether", "input unit")
	rootCmd.Flags().StringVarP(&rootCmdFlags.output, "output", "o", "wei", "output unit")
	rootCmd.Flags().StringVarP(&rootCmdFlags.amount, "amount", "a", "", "amount")

	rootCmd.Example = `
	ethconvert -i ether -o wei -a 1
	ethconvert 4.2 ether wei`
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
