package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var HTMLcomments *cobra.Command

func getHTMLcomments(cmd *cobra.Command, args []string) error {
	_url, _ := cmd.Flags().GetString("url")

	u, err := url.ParseRequestURI(_url)
	if err != nil {
		fmt.Print("Error while checking the URL\n")
		os.Exit(1)
	}

	response, err := http.Get(u.String())
	if err != nil {
		fmt.Print("Error while fetching the URL\n")
		os.Exit(1)
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print("Error while reading data from the URL\n")
		os.Exit(1)
	}

	re := regexp.MustCompile("<!--(.)*?-->")
	res := re.FindAllString(string(b), -1)
	if res == nil {
		fmt.Print("No HTML comments found.\n")
		os.Exit(1)
	}

	for _, m := range res {
		fmt.Println(m)
	}
	return nil
}

func init() {
	HTMLcomments = &cobra.Command{
		Use:   "htmlcom [URL]",
		Short: "Check if the web page contains HTML comments and extract them",
		RunE:  getHTMLcomments,
	}

	runCmd.AddCommand(HTMLcomments)
	HTMLcomments.PersistentFlags().String("url", "", "The website URL to scrap.")
	HTMLcomments.MarkPersistentFlagRequired("url")
}
