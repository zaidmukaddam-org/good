package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"good/helpers"

	"github.com/spf13/cobra"
)

var kanji *cobra.Command

var API_VERSION = "v1"

type Entry struct {
	Kanji         string
	Grade         int
	Meanings      []string
	Kun_readings  []string
	On_readings   []string
	Name_readings []string
	Jlpt          int
	Unicode       string
}

func runkanji(cmd *cobra.Command, args []string) error {
	kanji, _ := cmd.Flags().GetString("s")

	if helpers.IsKanji(kanji) == false {
		return errors.New("no kanji detected!")
	}

	url := fmt.Sprintf("https://kanjiapi.dev/%s/kanji/%s", API_VERSION, kanji)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read data from kanjiapi.dev: %s\n", err)
		os.Exit(1)
	}

	p := Entry{}
	err = json.Unmarshal(resBody, &p)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		os.Exit(1)
	}

	fmt.Printf("%s%s%s\n\n", helpers.RED, p.Kanji, helpers.END)
	if len(p.Meanings) != 0 {
		printInfo("meanings", strings.Join(p.Meanings, ", "))
	}
	if len(p.Kun_readings) != 0 {
		printInfo("kun readings", strings.Join(p.Kun_readings, ", "))
	}
	if len(p.On_readings) != 0 {
		printInfo("on readings", strings.Join(p.On_readings, ", "))
	}
	if len(p.Name_readings) != 0 {
		printInfo("name readings", strings.Join(p.Name_readings, ", "))
	}
	printInfo("jlpt", strconv.Itoa(p.Jlpt))
	printInfo("unicode", p.Unicode)
	printInfo("grade", strconv.Itoa(p.Grade))
	fmt.Print("\n")

	return nil
}

func printInfo(key string, text string) error {
	fmt.Print(helpers.SEP)
	fmt.Printf("%s: %s%s%s\n", key, helpers.CYAN, text, helpers.END)
	return nil
}

func init() {
	kanji = &cobra.Command{
		Use:   "kanji --s=[KANJI] (e.g., 猫)",
		Short: "Learn useful info about a specific Kanji",
		RunE:  runkanji,
	}
	funCmd.AddCommand(kanji)
	kanji.PersistentFlags().String("s", "", "The kanji character to search.")
	kanji.MarkPersistentFlagRequired("s")
}
