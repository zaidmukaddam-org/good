package cmd

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var memoCmd *cobra.Command

func runMemoCmd(cmd *cobra.Command, args []string) error {
	memory, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}

	fmt.Printf("%s\n", memory.String())

	return nil
}

func init() {
	memoCmd = &cobra.Command{
		Use:   "memory",
		Short: "Check your memory",
		RunE:  runMemoCmd,
	}
	checkCmd.AddCommand(memoCmd)
}
