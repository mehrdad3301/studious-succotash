package main 

import ( 
	"fmt"
	"flag"
	"strings"
	"net/http"
	"github.com/mehrdad3301/studious-succotash/4/link"
) 


func main() { 

	domain := flag.String("u", "https://www.calhoun.io", "domain") 
	flag.Parse() 

	urls, _ := getLinks(*domain)
	fmt.Println(urls)
	fmt.Println("_______________")
	urls = filterUrls(*domain, urls)
	fmt.Println(urls)
	
}


func getLinks(domain string) ([]string, error) { 

	resp, err := http.Get(domain)
	if err != nil { 
		return nil, err
	}

	links, err := link.Parse(resp.Body)
	if err != nil { 
		return nil, err
	}
	
	urls := make([]string, len(links))
	for i, link := range links { 
		urls[i] = link.Href
	}

	return urls, nil 
}


func filterUrls(domain string, urls []string) ([]string) { 

	var f []string 
	for _, u := range urls { 
		if strings.HasPrefix(u, "/") { 
			u = domain + u 
		}
		if strings.HasPrefix(u, domain) { 
			f = append(f, u)
		}
	}
	return f
}


//func bfs(domain string) map[string]struct{} { 
//
//	var queue []string 
//	visited := make(map[string]struct{})
//	
//	queue = append(queue, domain) 
//
//	for len(queue) != 0 { 
//
//		url := queue[0] 
//		queue = queue[1:] 
//		
//		visited[url] = nil 
//		
//		
//
//	}	
//
//	return struct{}
//
//} 

