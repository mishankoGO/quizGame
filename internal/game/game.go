package game

import (
	"context"
	"fmt"
	"github.com/mishankoGO/quizGame/internal/reader"
	"log"
)

func Game(ctx context.Context, reader reader.Reader) (int, int, error) {
	var correctCnt, totalCnt int

	for {
		record, err := reader.ReadLine()
		if err != nil {
			log.Println("error reading line: %v", err)
			return correctCnt, totalCnt, err
		}
		question, answer := record[0], record[1]
		fmt.Print(question + " ")

		ansCh := make(chan string)

		go waitForAnswer(ansCh)

		select {
		case <-ctx.Done():
			fmt.Println()
			return correctCnt, totalCnt, nil
		case resp := <-ansCh:
			if checkAnswer(resp, answer) {
				correctCnt++
			}
			totalCnt++
		}
	}
}

func waitForAnswer(ansCh chan string) {
	var response string
	_, err := fmt.Scanf("%s\n", &response)
	if err != nil {
		log.Println("error reading response")
	}
	ansCh <- response
}

func checkAnswer(response string, answer string) bool {
	return response == answer
}
