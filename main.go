package main

import (
	"context"
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

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), config.Limit)
	defer cancel()

	correctCnt, totalCnt, err := game.Game(ctx, r)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Time is up. Your result is: %d/%d", correctCnt, totalCnt)
}
