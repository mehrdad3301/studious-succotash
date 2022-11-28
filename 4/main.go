package main 

import ( 
	"fmt"
	"os"
	"flag"
	"github.com/mehrdad3301/studious-succotash/4/link"
)

func main() { 
	
	name := flag.String("n", "ex1.html", "html file name")
	flag.Parse()

	file, err := os.Open(*name)
	if err != nil { 
		fmt.Println(err)
	}
	links, _ := link.Parse(file)
	fmt.Println(links)
}
