package cmd

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var passwdCmd *cobra.Command

const (
	Letters = "abcdefghijklmnopqrstuvwxyz"
	Letter0 = "joli"
	Nums    = "0123456789"
	Num0    = "01"
	Symbols = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"
	Symbol0 = "<>[](){}:;'/|\\,"
)

func buildChars() string {
	var chars string

	chars += Letters
	chars = abracadabra(chars, Letter0)

	chars += strings.ToUpper(Letters)
	chars = abracadabra(chars, strings.ToUpper(Letter0))

	chars += Nums
	chars = abracadabra(chars, Num0)

	chars += Symbols
	chars = abracadabra(chars, Symbol0)

	return chars
}

func abracadabra(str, chars string) string {
	return strings.Map(func(r rune) rune {
		if !strings.ContainsRune(chars, r) {
			return r
		}
		return -1
	}, str)
}

func gegen(cmd *cobra.Command, args []string) error {
	var pwd string
	chars := strings.Split(buildChars(), "")
	max := big.NewInt(int64(len(chars)))
	l, _ := cmd.Flags().GetInt("l")
	L := uint64(l)

	if L < 16 || L > 64 {
		return errors.New("The length is not correct, give a value between 12 and 64, please.")
	}

	for i := uint64(0); i < L; i++ {
		val, err := rand.Int(rand.Reader, max)
		if err == nil {
			pwd += chars[val.Int64()]
		}
	}
	if pwd == "" {
		fmt.Println("Error when generating the password")
		os.Exit(1)
	}

	fmt.Printf("%s\n", pwd)
	return nil
}

func init() {
	passwdCmd = &cobra.Command{
		Use:   "passwordgen --l=[LENGTH]",
		Short: "Generate passwords",
		RunE:  gegen,
	}

	runCmd.AddCommand(passwdCmd)
	passwdCmd.PersistentFlags().Int("l", 16, "The length of the password (16-64 chars)")
}
