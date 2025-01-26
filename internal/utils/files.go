package utils

import (
	"encoding/json"
	"io"
	"os"
)

func ReadFile[T any](filename string) ([]T, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var items []T
	if len(data) > 0 {
		err = json.Unmarshal(data, &items)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func WriteFile[T any](filename string, list []T) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(list)
	if err != nil {
		return err
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}
