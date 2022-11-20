package main 

import ( 
	"os"
	"os/exec"
	"log"
	"fmt"
	"encoding/csv"
)

type problem struct { 
	question string 
	answer   string 
}

func main() { 
	
	file, err := os.Open("problems.csv")
	if err != nil { 
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil { 
		log.Fatal(err)
	}

	problems :=  getProblems(records)
	
	score := 0
	for i, problem := range problems { 

		clearScreen() 
		fmt.Printf("Q %d: %s? ", i, problem.question)
		
		var ans string 
		fmt.Scan(&ans)

		if ans == problem.answer{ 
			score += 1 
		}
	}
}


func clearScreen() { 
	cmd := exec.Command("clear")	
	cmd.Stdout = os.Stdout
	cmd.Run()	
}

func getProblems(records [][]string) []problem { 

	problems := make([]problem, len(records)) 
	for i, v := range records { 
		problems[i] = problem { v[0], v[1] } 		
	}	
	return problems
}

