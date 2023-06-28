package conf

import (
	"flag"
	"log"
	"os"
	"time"
)

// GameConfig is a config for server
type GameConfig struct {
	FileName string
	Limit    time.Duration
}

// InitFlags is responsible for parsing and setting flags for game
func InitFlags(conf *GameConfig) error {
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	// parse file name
	FileName := flags.String("f", conf.FileName, "file name")

	// parse time limit
	TimeLimit := flags.Duration("l", conf.Limit, "time limit")

	// parse the flags
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Println("error parsing flags: %v", err)
		return err
	}

	conf.FileName = *FileName
	conf.Limit = *TimeLimit
	return nil
}
