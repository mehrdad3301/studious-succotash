package main 

import ( 
	"fmt"
	"flag"
	"strings"
	"net/http"
	"github.com/mehrdad3301/studious-succotash/4/link"
) 


func main() { 

	domain := flag.String("u","https://www.calhoun.io", "domain") 
	flag.Parse() 
	
	m, err := bfs(*domain)
	if err != nil { 
		fmt.Println(err)
	}
	fmt.Printf("%+v", m)
	
}


func getLinks(domain string) ([]string, error) { 

	resp, err := http.Get(domain)
	if err != nil { 
		return nil, err
	}

	defer resp.Body.Close() 
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

		if strings.HasSuffix(u, "/") { 
			u = u[:len(u)-1]
		}

		if strings.HasPrefix(u, domain) { 
			f = append(f, u)
		}
	}
	return f
} 

func bfs(domain string) (map[string]bool, error) { 

	var queue []string 
	visited := make(map[string]bool)
	
	queue = append(queue, domain) 

	for len(queue) != 0 { 

		url := queue[0] 
		queue = queue[1:] 
		
		if visited[url] == true { 
			continue 
		}

		visited[url] = true 
		
		links, err := getLinks(url) 	
		if err != nil { 
			return nil, err
		}

		links = filterUrls(domain, links) 
		fmt.Println(url)
		
		for _, l := range links { 
			if _, ok := visited[l] ; !ok { 
				queue = append(queue, l)
			}
		}
	}	

	return visited, nil
} 

