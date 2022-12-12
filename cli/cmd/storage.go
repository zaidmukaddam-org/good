package cmd

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var storageCmd *cobra.Command

func runStorageCmd(cmd *cobra.Command, args []string) error {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v\n", err)
	}

	fmt.Printf("%v\n", block)

	for _, disk := range block.Disks {
		fmt.Printf(" %v\n", disk)
		for _, part := range disk.Partitions {
			fmt.Printf("  %v\n", part)
		}
	}
	return nil
}

func init() {
	storageCmd = &cobra.Command{
		Use:   "storage",
		Short: "Check your storage",
		RunE:  runStorageCmd,
	}
	checkCmd.AddCommand(storageCmd)
}
