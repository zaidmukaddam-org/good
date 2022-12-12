package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var osCmd *cobra.Command

func runOsCmd(cmd *cobra.Command, args []string) error {
	fmt.Printf("Your current OS is %s\n", runtime.GOOS)
	fmt.Printf("There are %d logical CPUs available\n", runtime.NumCPU())
	fmt.Printf("The root of the Go tree is in %s\n", runtime.GOROOT())
	return nil
}

func init() {
	osCmd = &cobra.Command{
		Use:   "os",
		Short: "Check your OS",
		RunE:  runOsCmd,
	}
	checkCmd.AddCommand(osCmd)
}
