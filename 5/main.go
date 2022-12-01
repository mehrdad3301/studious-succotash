package main 

import ( 
	"io"
	"os"
	"fmt"
	"flag"
	"strings"
	"net/url"
	"net/http"
	"encoding/xml"
	"github.com/mehrdad3301/studious-succotash/4/link"
) 


var xlmns = "https://www.sitemaps.org/schemas/sitemap/0.9/"

type loc struct { 
	Value string `xml:"loc"`
}
type urlset struct { 
	Urls  []loc `xml:"url"`
	Xlmns string `xml:"xlmns,attr"`
}

func main() { 

	url := flag.String("url","https://www.calhoun.io/", "url to build sitemap from") 
	depth := flag.Int("depth", 2, "depth limit for following links")
	flag.Parse() 

	links := bfs(*url, *depth)
	uSet := buildUrlSet(links)
	b, _:= xml.MarshalIndent(uSet, "", "\t")
	fmt.Fprintf(os.Stdout, xml.Header + string(b))
}

func buildUrlSet(links []string) urlset { 

	u := urlset{Urls: make([]loc, 0, len(links)), Xlmns: xlmns}	

	for _, l := range links { 
		u.Urls = append(u.Urls, loc{l})	
	}

	return u 
}


func get(domain string) []string { 

	resp, err := http.Get(domain)
	if err != nil { 
		return []string{}
	}

	defer resp.Body.Close() 

	reqURL := resp.Request.URL 
	baseURL := &url.URL { 
		Scheme: reqURL.Scheme, 
		Host :  reqURL.Host,
	}
	
	base := baseURL.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {

	links, err := link.Parse(r)
	if err != nil { 
		return []string{}
	}
	
	urls := make([]string, 0, len(links))
	for _, l := range links { 
		switch {
		case strings.HasPrefix(l.Href, "/"):
			urls = append(urls, base + l.Href)
		case strings.HasPrefix(l.Href, "http") :
			urls = append(urls, l.Href)
		}
	}
	return urls 
}

func filter(links []string, filterFn func(string) bool) ([]string) { 

	var ret []string 
	for _, u := range links { 
		if filterFn(u) { 
			ret = append(ret, u)
		}
	}
	return ret
} 


func withPrefix(prefix string) func(string) bool { 
	return func (link string) bool {
		return strings.HasPrefix(link, prefix) 
	}
}

func bfs(domain string, depth int) ([]string) { 

	seen := make(map[string]struct{})

	var q map[string]struct{} 
	nq := map[string]struct{}{ 
		domain: struct{}{}, 
	}

	for i := 0 ; i < depth ; i++ { 

		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 { 
			break 
		}

		for url, _ := range q { 

			if _, ok := seen[url] ; ok { 
				continue 
			}

			seen[url] = struct{}{}
			for _, link := range get(url) { 
				nq[link] = struct{}{}
			}
		}
		
	}

	ret := make([]string, 0, len(seen)) 
	for k, _ := range seen { 
		ret = append(ret, k)
	}	
	return ret 
} 

