package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

func newLink(href string, text string) *Link {
	return &Link{
		href: href,
		text: text,
	}
}

func (l Link) Print() {
	fmt.Printf("[Link] href:- %s text:- %s\n", l.href, l.text)
}

func readHTMLFile(filename string) (string, error) {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func getHref(node html.Token) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func parseHTMLlinks(text string) (data []Link) {
	tkn := html.NewTokenizer(strings.NewReader(text))
	var links []Link
	var isLink bool
	var href string
	var linkText string
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return links

		case tt == html.StartTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				href = getHref(t)
				isLink = true
			}

		case tt == html.TextToken:
			t := tkn.Token()
			if isLink {
				linkText += t.Data
			}

		case tt == html.EndTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				links = append(links, *newLink(href, strings.TrimSpace(linkText)))

				isLink = false
				linkText = ""
			}

		}

	}
	return links
}

func main() {
	data, err := readHTMLFile("./ex2.html")

	if err != nil {
		log.Fatal(err)
	}

	links := parseHTMLlinks(data)

	for _, link := range links {
		link.Print()
	}
}
