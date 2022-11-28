package link

import ( 
	"os"
	"io"
	"flag"
	"strings"
	"golang.org/x/net/html"
)

type Link struct { 
	Href string 
	Text string 
}

func Parse(r io.Reader) []Link{ 

}

func GetLinks(n *html.Node, links *[]Link) { 

	if n.Type == html.ElementNode { 
		if n.Data == "a" { 
			l := Link{ Href: getHref(n.Attr), 	
					   Text: getText(n)     }
			*links = append(*links, l)
			return  
		}
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling { 
		GetLinks(c, links)
	}
}

func getHref(attr []html.Attribute) string {
	for _, a := range attr { 
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
