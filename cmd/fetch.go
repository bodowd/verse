package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/html"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get verse",
	Run: func(cmd *cobra.Command, args []string) {

		input := strings.Join(args, " ")
		book, chapter, verseNum := parseQuery(input)

		books := booksInfo()
		bookMapping := books[book]

		url := buildURL(bookMapping.bookNumber, bookMapping.fullName, chapter)

		page := getHTML(url)

		doc, err := html.Parse(strings.NewReader(page))
		if err != nil {
			log.Fatal(err)
		}

		id := buildHTMLId(bookMapping.idAbbreviation, chapter, verseNum)

		tag := getElementById(doc, id)
		verse := getVerse(tag)
		finalVerse := strings.Join(verse, "")

		fmt.Println(finalVerse)
	},
}
