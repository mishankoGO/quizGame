package main

import (
	"log"
	"quizGame/conf"
	"quizGame/internal/reader/csvreader"
)

// variables to init
var (
	config conf.GameConfig // config with all the env variables
)

func main() {
	err := conf.InitFlags(&config)
	if err != nil {
		log.Println(err)
	}

	r := csvreader.NewCSVReader(config)
	defer r.Close()
	record, err := r.ReadLine()
	if err != nil {
		log.Println(err)
	}
	log.Println(record)
}
