package cmd

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"golang.org/x/net/html"
)

// helper functions
func readHtmlFromFile(fileName string) (string, error) {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func htmlPageToNode(fileName string) *html.Node {
	page, err := readHtmlFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(strings.NewReader(page))
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func testCaseToVerse(fileName string, id string) []string {
	doc := htmlPageToNode(fileName)
	element := traverse(doc, id)
	verse := getVerse(element)
	return verse
}

func assertStringsEqual(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

// tests
func TestTraverse(t *testing.T) {
	t.Run("successfully traverses to the id", func(t *testing.T) {
		doc := htmlPageToNode("psa_test_case.html")
		node := traverse(doc, "Psa1-1")
		got := renderNode(node)
		approvals.VerifyString(t, got)
	})
}

func TestGetVerse(t *testing.T) {
	t.Run("successfully gets verse with no extra elements", func(t *testing.T) {
		verse := testCaseToVerse("gen_test_case.html", "Gen1-1")
		got := cleanStrings(verse)
		want := "In the beginning God created the heavens and the earth."
		assertStringsEqual(t, got, want)
	})

	t.Run("successfully gets verse if there is an italics", func(t *testing.T) {
		verse := testCaseToVerse("italics_test_case.html", "Eph2-6")
		got := cleanStrings(verse)
		want := "And raised us up together with Him and seated us together with Him in the heavenlies in Christ Jesus,"
		assertStringsEqual(t, got, want)
	})

	t.Run("successfully gets verse if there is spans", func(t *testing.T) {
		doc := htmlPageToNode("span_test_case.html")
		element := traverse(doc, "Psa1-1")
		verse := getVerse(element)
		got := cleanStrings(verse)
		want := "Blessed is the man / Who does not walk / In the counsel of the wicked, / Nor stand on the path of sinners, / Nor sit in the seat of mockers;"
		assertStringsEqual(t, got, want)
	})
}
