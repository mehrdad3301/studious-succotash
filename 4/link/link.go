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

func Parse(r io.Reader) ([]Link, error){ 

	root, err := html.Parse(r)
	if err != nil { 
		return nil, err
	}

	linkNodes := getLinks(root)
	
	var links []Link
	for _, node := range linkNodes { 
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link { 
	return Link { Href: getHref(n), 
		   	      Text: getText(n) } 
}

func getLinks(n *html.Node) []*html.Node { 

	var links []*html.Node
	
	if n.Type == html.ElementNode && n.Data == "a" { 
		return []*html.Node{n} 
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling { 
		links = append(links, getLinks(c)...)
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

	var sb strings.Builder 
	if n.Type == html.TextNode { 
		return n.Data
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling { 
		sb.WriteString(getText(c))
	}
	return strings.Join(strings.Fields(sb.String()), " ")
}
