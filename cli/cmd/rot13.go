package cmd

import (
	"fmt"

	"good/helpers"

	"github.com/spf13/cobra"
)

var rot13Cmd *cobra.Command

func runRot13(cmd *cobra.Command, args []string) error {
	str, _ := cmd.Flags().GetString("str")
	helpers.RotReader(str)
	fmt.Print("\n\n")
	return nil
}

func init() {
	rot13Cmd = &cobra.Command{
		Use:   "rot13",
		Short: "encode/decode rot13 strings",
		RunE:  runRot13,
	}

	hackCmd.AddCommand(rot13Cmd)
	rot13Cmd.PersistentFlags().String("str", "", "The string to encode/decode")
	rot13Cmd.MarkPersistentFlagRequired("str")
}
