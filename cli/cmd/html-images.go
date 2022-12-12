package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"good/helpers"

	"github.com/spf13/cobra"
)

var HTMLimages *cobra.Command

const DEST = "good-images"

var (
	fileName string
	fileUrl  string
)

func getHTMLimages(cmd *cobra.Command, args []string) error {
	_url, _ := cmd.Flags().GetString("url")

	dest := fmt.Sprintf("%s/%s", helpers.GetHomedir(), DEST)

	if _, err := os.Stat(dest); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dest, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

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

	re := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	imgs := re.FindAllStringSubmatch(string(b), -1)
	out := make([]string, len(imgs))
	if out == nil {
		fmt.Print("No HTML images found.\n")
		os.Exit(1)
	}

	for i := range out {
		helpers.DownloadFile(imgs[i][1], dest)
	}

	return nil
}

func init() {
	HTMLimages = &cobra.Command{
		Use:   "htmlimg --url=[URL]",
		Short: "Download all images found on URL",
		RunE:  getHTMLimages,
	}

	runCmd.AddCommand(HTMLimages)
	HTMLimages.PersistentFlags().String("url", "", "The website URL to scrap.")
	HTMLimages.MarkPersistentFlagRequired("url")
}
