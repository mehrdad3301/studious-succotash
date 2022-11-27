package main 

import ( 
	_ "net/http"
	"encoding/json"
	"os"
	"fmt"
)


type Story struct { 

	Title 		string `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option`json:"options"`
}

/* 
	- marshalling 
	- template TODO 
	- http handlers TODO 	
*/

type Option struct { 

	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var stories = make(map[string]Story) 

func main() { 
	
	file, _ := os.ReadFile("gopher.json")
	json.Unmarshal(file, &stories)	
	
} 
