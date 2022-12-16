package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(booksCmd)
}

var booksCmd = &cobra.Command{
	Use:   "books",
	Short: "List of book abbreviations",
	Long:  "Information on book abbreviations to use in your query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Abbreviation to use in the command line app is on the left, and the book the abbreviation stands for is on the right")
		fmt.Println("usage: ")
		fmt.Println("$ verse Gen 1:1")
		fmt.Println(strings.Repeat("-", 80))
		booksInfo := booksInfo()
		var sortedBooks [][]string
		for k, v := range booksInfo {
			sortedBooks = append(sortedBooks, []string{k, v.fullName})
		}

		sort.Slice(sortedBooks, func(i, j int) bool {
			return sortedBooks[i][0] < sortedBooks[j][0]
		})

		for _, a := range sortedBooks {
			fmt.Printf("%s -- %s\n", a[0], a[1])
		}
	},
}
