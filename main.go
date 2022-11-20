package main 

import ( 
	"os"
	"log"
	"fmt"
	"time"
	"flag"
	"os/exec"
	"encoding/csv"
)

type problem struct { 
	question string 
	answer   string 
}

func main() { 
	
	timeLimit := flag.Int("t", 120, "quiz time limit")
	filename := flag.String("f", "problems.csv", "file path") 
	flag.Parse() 

	file, err := os.Open(*filename)
	if err != nil { 
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil { 
		log.Fatal(err)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	problems :=  getProblems(records)
	score := 0

	for i, problem := range problems { 

		clearScreen() 
		fmt.Printf("Q %d: %s? ", i, problem.question)

		answerCh := make(chan string)
		go func() {
				var ans string 
				fmt.Scan(&ans)
				answerCh <- ans
		}() 

		select { 
			case <-timer.C :
				fmt.Printf("you ran out of time!\n" + 
				"score %d out of %d", score, len(problems))
				return 
			case ans := <-answerCh :
				if ans == problem.answer { 
					score += 1 
				}			
		}
	}
	fmt.Printf("score %d out of %d", score, len(problems))
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

