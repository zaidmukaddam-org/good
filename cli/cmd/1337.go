package cmd

import (
	"fmt"

	"good/helpers"

	"github.com/spf13/cobra"
)

var leet *cobra.Command

func convert2leet(cmd *cobra.Command, args []string) error {
	t, _ := cmd.Flags().GetString("t")
	fmt.Printf("%s\n\n", helpers.Str2leet(t))
	return nil
}

func init() {
	leet = &cobra.Command{
		Use:   "1337 --t=[TEXT]",
		Short: "Convert a string to leetspeak and shine on Internet",
		RunE:  convert2leet,
	}
	funCmd.AddCommand(leet)
	leet.PersistentFlags().String("t", "", "The text to convert")
	leet.MarkPersistentFlagRequired("t")
}
