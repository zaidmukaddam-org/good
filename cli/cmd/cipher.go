package cmd

import (
	"errors"
	"fmt"

	"good/helpers"

	"github.com/spf13/cobra"
)

var cipherCmd *cobra.Command

func encipher(key string, text string) error {
	t := helpers.BrokenCipher(key, text)
	fmt.Printf("Result:\n\n %s\n\n", t)
	return nil
}

func decipher(key string, cipher string) error {
	t := helpers.BrokenDeCipher(key, cipher)
	fmt.Printf("Result:\n\n %s\n\n", t)
	return nil
}

func runCipherCmd(cmd *cobra.Command, args []string) error {
	message, _ := cmd.Flags().GetString("m")
	key, _ := cmd.Flags().GetString("k")
	d, _ := cmd.Flags().GetBool("decipher")

	if len(key) < 12 {
		return errors.New("The key must contain at least 12 chars!")
	}

	if d {
		decipher(key, message)
	} else {
		encipher(key, message)
	}

	return nil
}

func init() {
	cipherCmd = &cobra.Command{
		Use:   "cipher --m=[MESSAGE] --k=[KEY] --decipher",
		Short: "Encipher or decipher your secret message",
		RunE:  runCipherCmd,
	}

	hackCmd.AddCommand(cipherCmd)
	cipherCmd.PersistentFlags().String("m", "", "The message to encipher")
	cipherCmd.PersistentFlags().String("k", "", "The key for the operation")
	cipherCmd.PersistentFlags().Bool("decipher", false, "Decipher a cipher message")
	cipherCmd.MarkPersistentFlagRequired("m")
	cipherCmd.MarkPersistentFlagRequired("k")
}
