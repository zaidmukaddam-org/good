package cmd

import (
	"fmt"
	"os"

	"good/helpers"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "good",
	Short: "The root command. Please, check subcommands below.",
	Long:  "Personal playground composed of CLI commands made in go.",
}

var funCmd = &cobra.Command{
	Use:   "fun",
	Short: "commands for fun",
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "helpful commands to check your system",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run some analysis & use generators",
}

var hackCmd = &cobra.Command{
	Use:   "hack",
	Short: "commands for hackers",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	helpers.Header()

	rootCmd.AddCommand(funCmd)
	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(hackCmd)

}
