package avito

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/url"
	"path"
)

const baseUrl = "https://avito.ru/"

func prepareUrlWithPath(urlPath string) *url.URL {
	reqUrl, _ := url.Parse(baseUrl)
	reqUrl.Path = path.Join(reqUrl.Path, urlPath)
	return reqUrl
}

func firstImmediateText(s *goquery.Selection) string {

	for _, node := range s.Nodes {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if child.Type == html.TextNode {
				return child.Data
			}
		}
	}

	return ""
}