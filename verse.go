package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHTML(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)

}

func getAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func checkId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := getAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string) *html.Node {
	if checkId(n, id) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := traverse(c, id)
		if res != nil {
			return res
		}
	}
	return nil
}

func getElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id)
}

func getVerse(n *html.Node) []string {
	fmt.Println(renderNode(n))
	var verse []string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(c.Data)
		if c.Data == "b" {
			verse = append(verse, c.NextSibling.Data)
		} else if c.Data == "i" {
			// sometimes there's an italics in the middle of the verse
			// get the data in between the italics tag with FirstChild
			// then continue on to get the rest of the data in b tag, which is
			// in NextSibling.Data

			verse = append(verse, c.FirstChild.Data)
			// append the rest of the b tag
			verse = append(verse, c.NextSibling.Data)
		}

	}

	return verse

}

func renderNode(n *html.Node) string {

	var buf bytes.Buffer
	w := io.Writer(&buf)

	err := html.Render(w, n)

	if err != nil {
		return ""
	}

	return buf.String()
}

func parseQuery(query string) (string, string, string) {
	var book strings.Builder
	var chapter strings.Builder
	var verse strings.Builder

	sepCount := 0
	for i := 0; i < len(query); i++ {
		if sepCount == 0 {
			if string(query[i]) == " " {
				sepCount++
				continue
			}
			book.WriteString(string(query[i]))
		} else if sepCount == 1 {
			if string(query[i]) == ":" {
				sepCount++
				continue
			}
			chapter.WriteString(string(query[i]))
		} else {
			verse.WriteString(string(query[i]))
		}

	}
	return book.String(), chapter.String(), verse.String()
}

func main() {

	input := "John 3:16"
	book, chapter, verseNum := parseQuery(input)
	fmt.Printf("book %s, chapter %s, verse %s", book, chapter, verseNum)

	// url := "https://text.recoveryversion.bible/01_Genesis_1.htm"
	// url := "https://text.recoveryversion.bible/11_1Kings_1.htm"
	// url := "https://text.recoveryversion.bible/11_1Kings_8.htm"
	url := "https://text.recoveryversion.bible/62_1John_1.htm"
	// bookNumber_Book_Chapter.htm
	page := getHTML(url)

	doc, err := html.Parse(strings.NewReader(page))

	if err != nil {
		log.Fatal(err)
	}

	tag := getElementById(doc, "FJo1-2")
	verse := getVerse(tag)
	fmt.Println(verse)

	bookNumbers := getBookNumbers()
	fmt.Println(bookNumbers)
}
