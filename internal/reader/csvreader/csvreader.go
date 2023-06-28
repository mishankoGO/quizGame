package csvreader

import (
	"encoding/csv"
	"github.com/mishankoGO/quizGame/conf"
	"github.com/mishankoGO/quizGame/internal/reader"
	"log"
	"os"
)

type CSVReader struct {
	config conf.GameConfig
	reader csv.Reader
	file   os.File
}

func NewCSVReader(config conf.GameConfig) reader.Reader {
	file, err := os.Open(config.FileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)

	return &CSVReader{
		reader: *csvReader,
		config: config,
		file:   *file,
	}
}

func (r *CSVReader) Close() error {
	err := r.file.Close()
	if err != nil {
		log.Printf("error closing file: %v", err)
		return err
	}
	return nil
}

func (r *CSVReader) ReadLine() ([]string, error) {
	record, err := r.reader.Read()
	if err != nil {
		log.Printf("error reading a line: %v", err)
		return nil, err
	}
	return record, err
}
