package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Problem store the problem information
type Problem struct {
	question string
	answer   string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "Archivo csv con el formato: 'question,answer'")
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

	problems := parseCSVProblems(csvProblems)

	corrects := 0
	var answer string
	for index, problem := range problems {
		fmt.Printf("[%d] %s = ", index+1, problem.question)
		fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			corrects++
		}
	}

	fmt.Printf("Tu puntaje fué: %d de: %d\n", corrects, len(problems))

}

func parseCSVProblems(csvProblems [][]string) []Problem {

	problems := make([]Problem, len(csvProblems))

	for i := 0; i < len(csvProblems); i++ {
		problems[i] = Problem{
			question: csvProblems[i][0],
			answer:   strings.TrimSpace(csvProblems[i][1]),
		}
	}
	return problems

}
