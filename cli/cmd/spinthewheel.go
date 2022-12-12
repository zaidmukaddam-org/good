package cmd

import (
	"fmt"
    "math/rand"
	"time"

	"github.com/spf13/cobra"
)

var decisions = []string {
	"Do it. It's a no-brainer!",
	"Don't do it. It's a trap!",
	"Hum... maybe talk about it with your friends and family before doing something silly...",
	"It's too soon to turn the wheel. Try again next week, please.",
	"The Oracle said it's time because the moon told him so, but it won't be easy.",
	"The chance don't always smile to those who dare, but only those who never perform make no mistake.",
}

var spinTheWheelCmd *cobra.Command

func runSpinTheWheelCmd(cmd *cobra.Command, args []string) error {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("%s\n\n", decisions[rand.Intn(5)])
	return nil
}

func init() {
	spinTheWheelCmd = &cobra.Command{
		Use:   "spinthewheel",
		Short: "Don't know what to decide? I can't guarantee anything, but maybe leave it to chance...",
		RunE:  runSpinTheWheelCmd,
	}

	funCmd.AddCommand(spinTheWheelCmd)
}
