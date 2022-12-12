package cmd

import (
	"fmt"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var gpuCmd *cobra.Command

func runGpuCmd(cmd *cobra.Command, args []string) error {
	gpu, err := ghw.GPU()
	if err != nil {
		fmt.Printf("Error getting GPU info: %v\n", err)
	}

	fmt.Printf("%v\n", gpu)

	for _, card := range gpu.GraphicsCards {
		fmt.Printf(" %v\n", card)
	}
	return nil
}

func init() {
	gpuCmd = &cobra.Command{
		Use:   "gpu",
		Short: "Check your GPU",
		RunE:  runGpuCmd,
	}
	checkCmd.AddCommand(gpuCmd)
}
