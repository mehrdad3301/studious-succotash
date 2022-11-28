package main 

import ( 
	"fmt"
	"flag"
	"github.com/mehrdad3301/studious-succotash/4/link"
)

func main() { 
	
	name := flag.String("n", "ex1.html", "html file name")
	flag.Parse()

	file, _:= os.Open(*name)
	if err != nil { 
		fmt.Println(err)
	}
	link.Parse(file)
}
