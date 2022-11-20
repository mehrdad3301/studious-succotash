package main 

import ( 
	"os"
	"os/exec"
	"log"
	"fmt"
	"encoding/csv"
)

var ( 
	q_number = 0 
	score = 0
	ans string
)

func main() { 
	
	file, err := os.Open("problems.csv")
	if err != nil { 
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	record, _ := reader.Read()
	for record != nil{ 
		clearScreen() 
		askQuestion(record[0], record[1])
		record, _ = reader.Read()
	}
	fmt.Printf("you scored %d", score) 
}


func clearScreen() { 
	cmd := exec.Command("clear")	
	cmd.Stdout = os.Stdout
	cmd.Run()	
}

func askQuestion(question, answer string) { 

		fmt.Printf("question: %d \n what is %s? ", q_number, question)

		fmt.Scan(&ans)
		if ans == answer { 
			score += 1 
		}
		q_number += 1 
}
