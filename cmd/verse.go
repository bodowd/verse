package cmd

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
	if n == nil {
		return []string{"Verse cannot be found."}
	}
	// starts at node n which is the <p> containing the id we are interested in
	var verse []string

	// FirstChild of n is <b>
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "b" {
			// the FirstChild of <b> is the <a> tags
			// the Next sibling of <b> could contain a verse or <span>
			// the c.NextSibling.NextSibling could contain "i" which we check
			// in the next else if
			if c.NextSibling.Data != "span" {
				verse = append(verse, c.NextSibling.Data)
			}
		} else if c.Data == "i" {
			// sometimes there's an italics in the middle of the verse
			// get the word inbetween the <i> tag is the First Child
			// then continue on to get the rest of the data in b tag, which is
			// in NextSibling.Data
			verse = append(verse, c.FirstChild.Data)
			// append the rest of the b tag
			verse = append(verse, c.NextSibling.Data)
			// then the next iteration will happen which will go to c.NextSibling
			// which could be another <i> tag which we will then go back into
			// this code block. This will end when c.NextSibling.NextSibling == nil
			// which will end the for loop
		} else if c.Data == "span" {
			verse = traverseNestedSpans(c, verse)
		}
	}
	return verse
}

// the FirstChild will contain the first block of text in the first
// <span>
//
// The next <span> blocks are Nextsiblings of FirstChild
// i.e. c.FirstChild.NextSibling == <span>
//
// so c.FirstChild.NextSibling.FirstChild contains the next text
// within that next <span>
//
// i.e. c.FirstChild.NextSibling.FirstChild == text
//
// the pattern seems to be FirstChild.NextSibling => text in one <span>
// then FirstChild.NextSibling.FirstChild.NextSibling => the next
// text in the following <span>
//
// the loop needs to be different for the <span>.
//
// in the first hit of "span", it gives the whole block.
// there is no NextSibling. I.e. c.NextSibling == nil
//
// So from the text of this current node (the FirstChild of the current <span>)
// we need to go to the NextSibling of the text (the NextSibling of FirstChild)
// this will contain the next <span>
//
//	<span current>
//		First Child -- text
//		    <span Next Sibling>
//		        First Child -- text
//		        <span NextSibling>
//		            First Child -- text
//		        </span>
//		    </span></span>
func traverseNestedSpans(c *html.Node, verse []string) []string {

	if c.Data == "span" {
		// append text that is under the span
		verse = append(verse, c.FirstChild.Data)

	}

	// then traverse nested elements
	for c := c.FirstChild.NextSibling; c != nil; c = c.FirstChild.NextSibling {

		verse = append(verse, c.FirstChild.Data)
		if c.Data == "i" {
			// if the verse ends with an italics tag, there is no NextSibling
			if c.NextSibling != nil {
				// append the rest of the span tag similar to how we handle
				// italics in the getVerse if "i" conditional
				verse = append(verse, c.NextSibling.Data)
				if c.NextSibling.NextSibling != nil {
					//  Then traverse the next span
					verse = traverseNestedSpans(c.NextSibling.NextSibling, verse)
				}
			}
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

func buildURL(bookNumber, bookName, chapter string) string {
	// ...bookNumber_fullName_Chapter.htm
	URL := fmt.Sprintf("https://text.recoveryversion.bible/%s_%s_%s.htm", bookNumber, bookName, chapter)

	return URL
}

func buildHTMLId(idAbbreviation, chapter, verseNum string) string {
	return idAbbreviation + chapter + "-" + verseNum
}
