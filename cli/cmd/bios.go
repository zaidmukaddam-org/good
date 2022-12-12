package cmd

import (
	"fmt"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var biosCmd *cobra.Command

func runBiosCmd(cmd *cobra.Command, args []string) error {
	bios, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("error getting BIOS info: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", bios.YAMLString())
	return nil
}

func init() {
	biosCmd = &cobra.Command{
		Use:   "bios",
		Short: "Check your bios",
		RunE:  runBiosCmd,
	}
	checkCmd.AddCommand(biosCmd)
}
