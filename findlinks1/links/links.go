package links

import (
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)

				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)

	return links, nil
}
