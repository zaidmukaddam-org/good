package cmd

import (
	"fmt"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/spf13/cobra"
)

var chassisCmd *cobra.Command

func runChassisCmd(cmd *cobra.Command, args []string) error {
	chassis, err := ghw.Chassis()
	if err != nil {
		fmt.Printf("error getting chassis info: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", chassis.YAMLString())
	return nil
}

func init() {
	chassisCmd = &cobra.Command{
		Use:   "chassis",
		Short: "Check your chassis info",
		RunE:  runChassisCmd,
	}
	checkCmd.AddCommand(chassisCmd)
}
