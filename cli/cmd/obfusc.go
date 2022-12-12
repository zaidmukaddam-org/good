package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var obfCmd *cobra.Command

func runObfCmd(cmd *cobra.Command, args []string) error {
	str, _ := cmd.Flags().GetString("str")
	deob, _ := cmd.Flags().GetBool("reverse")

	if deob {
		fmt.Printf("%s\n\n", decodeHex(str))
	} else {
		fmt.Printf("%s\n\n", encodeHex(str))
	}

	return nil
}

func encodeHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

func decodeHex(str string) string {
	d, err := hex.DecodeString(str)
	if err != nil {
		fmt.Printf("Error when decoding the string: %s\n\n", err)
		os.Exit(1)
	}
	return string(d)
}

func init() {
	obfCmd = &cobra.Command{
		Use:   "obfuscate --str=[STR]",
		Short: "obfuscate/deobfuscate string with hexadecimal encoding",
		RunE:  runObfCmd,
	}

	hackCmd.AddCommand(obfCmd)
	obfCmd.PersistentFlags().String("str", "", "The string to obfuscate")
	obfCmd.PersistentFlags().Bool("reverse", false, "Deobfuscate instead of obfustating")
	obfCmd.MarkPersistentFlagRequired("str")
}
