package link

import (
	"io"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseFile(content io.Reader) ([]Link, error) {
	node, err := html.Parse(content)
	if err != nil {
		return nil, err
	}
	return parseNode(node)
}

func buildLink(node *html.Node) Link {
	text := getNodeText(node)
	href := getHref(node)
	return Link{Href: href, Text: strings.TrimSpace(text)}
}

func getHref(node *html.Node) string {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func parseNode(node *html.Node) ([]Link, error) {
	var links []Link = []Link{}
	if node.Type == html.ElementNode && node.Data == "a" {
		links = append(links, buildLink(node))
	} else {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			childLinks, err := parseNode(child)
			if err != nil {
				return nil, err
			}
			links = append(links, childLinks...)
		}
	}
	return links, nil
}

var spaceRe = regexp.MustCompile(`\s+`)

func getNodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		return spaceRe.ReplaceAllString(node.Data, " ")
	}
	text := ""
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		text += getNodeText(child)
	}
	return text
}
