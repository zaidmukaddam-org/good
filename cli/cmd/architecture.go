package cmd

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var archiCmd *cobra.Command

func runArchiCmd(cmd *cobra.Command, args []string) error {
	topology, err := ghw.Topology()
	if err != nil {
		fmt.Printf("Error getting topology info: %v\n", err)
	}

	fmt.Printf("%v\n", topology)

	for _, node := range topology.Nodes {
		fmt.Printf(" %v\n", node)
		for _, cache := range node.Caches {
			fmt.Printf("  %v\n", cache)
		}
	}
	return nil
}

func init() {
	archiCmd = &cobra.Command{
		Use:   "archi",
		Short: "Check your architecture with the Topology function",
		RunE:  runArchiCmd,
	}
	checkCmd.AddCommand(archiCmd)
}
