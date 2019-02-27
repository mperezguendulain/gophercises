package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Problem store the problem information
type Problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "Archivo csv con el formato: 'question,answer'")
	limitToQuiz := flag.Int("limit", 30, "Limite de tiempo para el quiz")
	shuffleQuiz := flag.Bool("shuffle", false, "Indica si se barajéan las preguntas para que estás se muestren aleatoriamente")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo: ", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	csvProblems, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo: ", *csvFileName)
		os.Exit(1)
	}

	problems := parseCSVProblems(csvProblems, *shuffleQuiz)

	timer := time.NewTimer(time.Duration(*limitToQuiz) * time.Second)

	corrects := 0
	var answer string
	answerChan := make(chan string)
	for index, problem := range problems {
		fmt.Printf("[%d] %s = ", index+1, problem.question)

		go func() {
			fmt.Scanf("%s", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nTu puntaje fué: %d de: %d\n", corrects, len(problems))
			return
		case answer := <-answerChan:
			if answer == problem.answer {
				corrects++
			}
		}

	}

	fmt.Printf("Tu puntaje fué: %d de: %d\n", corrects, len(problems))

}

func parseCSVProblems(csvProblems [][]string, shuffle bool) []Problem {

	problems := make([]Problem, len(csvProblems))

	for i := 0; i < len(csvProblems); i++ {
		problems[i] = Problem{
			question: csvProblems[i][0],
			answer:   strings.TrimSpace(csvProblems[i][1]),
		}
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	return problems

}
