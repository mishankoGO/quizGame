package game

import (
	"fmt"
	"github.com/mishankoGO/quizGame/internal/reader"
	"log"
)

func Game(reader reader.Reader) (int, int, error) {
	var correctCnt, totalCnt int

	for {
		record, err := reader.ReadLine()
		if err != nil {
			log.Println("error reading line: %v", err)
			return correctCnt, totalCnt, err
		}

		if len(record) != 0 {
			question, answer := record[0], record[1]

			var response string
			fmt.Print(question + " ")
			_, err = fmt.Scanf("%s", &response)
			if err != nil {
				log.Println("error reading response")
				return correctCnt, totalCnt, err
			}

			if checkAnswer(response, answer) {
				correctCnt++
			}
			totalCnt++
		} else {
			return correctCnt, totalCnt, nil
		}
	}
}

func checkAnswer(response string, answer string) bool {
	return response == answer
}
