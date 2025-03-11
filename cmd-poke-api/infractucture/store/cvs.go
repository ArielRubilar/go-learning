package store

import (
	"os"

	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

type CvsStore struct {
}

func NewCvsStore() repository.Store {
	return &CvsStore{}
}

func (s *CvsStore) Save(key string, data []map[string]string) error {

	os.Mkdir("data", 0777)

	file, err := os.Create("data/" + key + ".csv")

	if data[0] == nil {
		return nil
	}

	header := ""

	for h := range data[0] {
		header += h + ","
	}

	file.WriteString(header + "\n")

	for _, row := range data {
		values := ""
		for _, value := range row {
			values += value + ","
		}
		file.WriteString(values + "\n")
	}

	if err != nil {
		return err
	}

	return nil
}
