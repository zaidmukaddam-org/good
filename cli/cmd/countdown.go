package cmd

import (
	"fmt"
	"time"

	"good/helpers"

	"github.com/spf13/cobra"
)

var countDownCmd *cobra.Command

const OFFSET = 2

func rungCountDownCmd(cmd *cobra.Command, args []string) error {
	d, _ := cmd.Flags().GetInt("d")
	tick := time.Tick(1 * time.Second)
	duration := int64(d + OFFSET)
	end := time.After(time.Duration(duration) * time.Second)
	c := d

	fmt.Printf("%s%d seconds left!%s\n\n\n", helpers.RED, c+OFFSET, helpers.END)
	time.Sleep(OFFSET * time.Second)

	for i := 0; i < 3; i++ {
		printSep()
	}

	for {
		select {
		case <-tick:
			if c < 6 {
				fmt.Printf("%s%d%s\n", helpers.RED, c, helpers.END)
			} else {
				fmt.Println(c)
			}
			c--
		case <-end:
			fmt.Print("\n\nBim! It's over.\n\n")
			return nil
		default:
			printSep()
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func printSep() {
	fmt.Println("    ")
}

func init() {
	countDownCmd = &cobra.Command{
		Use:   "countdown --d=[DURATION] (in seconds)",
		Short: "Countdown",
		RunE:  rungCountDownCmd,
	}

	runCmd.AddCommand(countDownCmd)
	countDownCmd.PersistentFlags().Int("d", 10, "The duration")
}
