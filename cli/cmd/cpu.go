package cmd

import (
	"fmt"
	"math"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var cpuCmd *cobra.Command

func runCpuCmd(cmd *cobra.Command, args []string) error {
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v\n", err)
	}

	fmt.Printf("%v\n", cpu)

	for _, proc := range cpu.Processors {
		fmt.Printf(" %v\n", proc)
		for _, core := range proc.Cores {
			fmt.Printf("  %v\n", core)
		}
		if len(proc.Capabilities) > 0 {
			rows := int(math.Ceil(float64(len(proc.Capabilities)) / float64(6)))
			for row := 1; row < rows; row = row + 1 {
				rowStart := (row * 6) - 1
				rowEnd := int(math.Min(float64(rowStart+6), float64(len(proc.Capabilities))))
				rowElems := proc.Capabilities[rowStart:rowEnd]
				capStr := strings.Join(rowElems, " ")
				if row == 1 {
					fmt.Printf("  capabilities: [%s\n", capStr)
				} else if rowEnd < len(proc.Capabilities) {
					fmt.Printf("                 %s\n", capStr)
				} else {
					fmt.Printf("                 %s]\n", capStr)
				}
			}
		}
	}

	return nil
}

func init() {
	cpuCmd = &cobra.Command{
		Use:   "cpu",
		Short: "Check your CPU",
		RunE:  runCpuCmd,
	}
	checkCmd.AddCommand(cpuCmd)
}
