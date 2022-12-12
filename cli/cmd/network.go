package cmd

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var networkCmd *cobra.Command

func runNetworkCmd(cmd *cobra.Command, args []string) error {
	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting network info: %v", err)
	}

	fmt.Printf("%v\n", net)

	for _, nic := range net.NICs {
		fmt.Printf(" %v\n", nic)

		enabledCaps := make([]int, 0)
		for x, cap := range nic.Capabilities {
			if cap.IsEnabled {
				enabledCaps = append(enabledCaps, x)
			}
		}
		if len(enabledCaps) > 0 {
			fmt.Printf("  enabled capabilities:\n")
			for _, x := range enabledCaps {
				fmt.Printf("   - %s\n", nic.Capabilities[x].Name)
			}
		}
	}

	return nil
}

func init() {
	networkCmd = &cobra.Command{
		Use:   "network",
		Short: "Check your network",
		RunE:  runNetworkCmd,
	}
	checkCmd.AddCommand(networkCmd)
}
