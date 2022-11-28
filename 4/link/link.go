package link

import ( 
	"io"
	"strings"
	"golang.org/x/net/html"
)

type Link struct { 
	Href string 
	Text string 
}

func Parse(r io.Reader) []Link{ 

	root, err := html.Parse(r)
	if err != nil { 
		panic(err)
	}

	linkNodes = getLinks(root)
	
	var links []Link
	for _, node := range linkNodes { 
		links = append(links, buildLink(node))
	}
	return links
}

func buildLink(n *html.Node) Link { 
	return Link { Href: getHref(n), 
		   	      Text: getText(n) } 
}

func getLinks(n *html.Node) []*html.Node { 

	var links []*html.Node
	
	if n.Type == html.ElementNode && n.Data == "a" { 
		return  n 
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling { 
		links = append(links, getLinks(c))
	}

	return links
}

func getHref(n *html.Node) string {
	for _, a := range n.Attr { 
		if a.Key == "href" { 
			return a.Val
		}
	}
	return ""
}

func getText(n *html.Node) string { 

	var texts []string

	if n.Type == html.TextNode { 
		return n.Data
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling { 
		texts = append(texts, getText(c))
	}
	return strings.Join(texts, "")
}
