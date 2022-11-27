package main 

import ( 
	"net/http"
	"html/template"
	"encoding/json"
	"os"
)


type Story struct { 

	Title 		string `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option`json:"options"`
}

type Option struct { 

	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var (

	stories map[string]Story 
	storyTemp *template.Template
)
	
func init() { 
	stories = make(map[string]Story) 
	storyTemp, _ = template.ParseFiles("temp.html")
}

func main() { 
	
	file, _ := os.ReadFile("gopher.json")
	json.Unmarshal(file, &stories)	
	http.HandleFunc("/", handler)
	
	http.ListenAndServe("localhost:8080", nil) 
} 

func handler(w http.ResponseWriter, r *http.Request) { 
	path := r.URL.Path[1:]
	if path != "" { 
		storyTemp.Execute(w, stories[path])
	} else {
		http.Redirect(w, r, "/intro", http.StatusFound)
	}
}
