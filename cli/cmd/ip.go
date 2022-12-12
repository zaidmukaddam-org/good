package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var ipCmd *cobra.Command

type Data struct {
	Ip string
}

func runIpCmd(cmd *cobra.Command, args []string) error {
	res, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read your public IP from ipinfo.io: %s\n", err)
		os.Exit(1)
	}

	var payload Data
	err = json.Unmarshal(resBody, &payload)
	if err != nil {
		fmt.Printf("Error during Unmarshal(): %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Your public IP is %s\n", payload.Ip)

	return nil
}

func init() {
	ipCmd = &cobra.Command{
		Use:   "ip",
		Short: "Check your public IP",
		RunE:  runIpCmd,
	}
	checkCmd.AddCommand(ipCmd)
}
