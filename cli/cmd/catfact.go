package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var catFactCmd *cobra.Command

type CatFact struct {
	Fact string
}

func runCatFactCmd(cmd *cobra.Command, args []string) error {
	res, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read data from catfact.ninja: %s\n", err)
		os.Exit(1)
	}

	var payload CatFact
	err = json.Unmarshal(resBody, &payload)
	if err != nil {
		fmt.Printf("Error during Unmarshal(): %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n\n", payload.Fact)

	return nil
}

func init() {
	catFactCmd = &cobra.Command{
		Use:   "catfact",
		Short: "Get random cat fact",
		RunE:  runCatFactCmd,
	}
	funCmd.AddCommand(catFactCmd)
}
