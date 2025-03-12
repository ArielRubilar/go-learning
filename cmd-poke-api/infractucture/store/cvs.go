package store

import (
	"os"
	"strings"

	repository "github.com/ArielRubilar/go-learning/domain/repositories"
)

type CvsStore struct {
}

func NewCvsStore() repository.Store {
	return &CvsStore{}
}

func toCsv(data []string) string {
	csv := ""

	for index, d := range data {
		csv += d
		if index < len(data)-1 {
			csv += ","
		}
	}

	return csv
}

func createDir(key string) {
	folders := strings.Split(key, "/")

	if len(folders) > 1 {
		fullDirPath := strings.Join(folders[:len(folders)-1], "/")
		os.MkdirAll(fullDirPath, os.ModePerm)
	}
}

func (s *CvsStore) Save(key string, data []map[string]string) error {

	createDir(key)

	file, err := os.Create(key + ".csv")

	if data[0] == nil {
		return nil
	}

	headers := []string{}

	for h := range data[0] {
		headers = append(headers, h)
	}

	file.WriteString(toCsv(headers) + "\n")

	for _, row := range data {
		values := []string{}
		for _, head := range headers {
			values = append(values, row[head])
		}
		file.WriteString(toCsv(values) + "\n")
	}

	if err != nil {
		return err
	}

	return nil
}
