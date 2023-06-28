package main

import (
	"fmt"
	"github.com/mishankoGO/quizGame/conf"
	"github.com/mishankoGO/quizGame/internal/game"
	"github.com/mishankoGO/quizGame/internal/reader/csvreader"
	"log"
)

// variables to init
var (
	config conf.GameConfig // config with all the env variables
)

func main() {
	// init config
	err := conf.InitFlags(&config)
	if err != nil {
		log.Println(err)
	}

	// init reader
	r := csvreader.NewCSVReader(config)
	defer r.Close()

	correctCnt, totalCnt, err := game.Game(r)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Your result: %d/%d", correctCnt, totalCnt)
}
