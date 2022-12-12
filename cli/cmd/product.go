package cmd

import (
	"fmt"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var productCmd *cobra.Command

func runProductCmd(cmd *cobra.Command, args []string) error {
	product, err := ghw.Product()
	if err != nil {
		fmt.Printf("error getting product info: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", product.YAMLString())
	return nil
}

func init() {
	productCmd = &cobra.Command{
		Use:   "product",
		Short: "Check your product info (some info might require root access)",
		RunE:  runProductCmd,
	}
	checkCmd.AddCommand(productCmd)
}
