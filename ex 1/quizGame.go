package main

import (
	"encoding/csv"
	"flag"
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
	var path *string = flag.String("path", "./problems.csv", "path")
	var seconds *int = flag.Int("seconds", 10, "seconds")
	flag.Parse()
	questions := readcsv(*path)
	defer result(&correctAnswers, len(questions))
	fmt.Println("Press the Enter Key to Start")
	fmt.Scanln()
	go askQuestions(&questions, &correctAnswers)
	time.Sleep(time.Duration(*seconds) * time.Second)
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
	defer recordFile.Close()
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
