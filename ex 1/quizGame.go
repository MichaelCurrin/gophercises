package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Question struct {
	quest  string
	answer string
}

func main() {
	var correctAnswers int = 0
	fmt.Println("hello world")
	questions := readcsv("./problems.csv")
	defer result(&correctAnswers, len(questions))
	fmt.Println("Press the Any Key to Start")
	fmt.Scanln()
	go askQuestions(&questions, &correctAnswers)
	time.Sleep(10 * time.Second)
}

func askQuestions(questions *[]Question, correctAnswers *int) {
	var answer string
	for i, quest := range *questions {
		fmt.Println("Question", i+1, ":", quest.quest)
		fmt.Scanln(&answer)
		if quest.answer == answer {
			*correctAnswers++
		}
	}
}
func result(correctAnswers *int, quantityQuestions int) {
	fmt.Println("Number of correct answers: ", *correctAnswers)
	fmt.Println("Number of total answers: ", quantityQuestions)
}
func readcsv(path string) []Question {
	var quests []Question
	recordFile, err := os.Open(path)
	if err != nil {
		fmt.Println("An error encounteres ::", err)
	}

	reader := csv.NewReader(recordFile)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		quests = append(quests, Question{quest: record[0], answer: record[1]})
	}
	return quests
}
