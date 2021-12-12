package asciimodel

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type Model struct {
	file string
	data map[string]string
}

func (model *Model) Load() error {
	file, err := ioutil.ReadFile("./data/data.json")
	if err != nil {
		log.Fatal("Error on init of asciibulder, trying to read data.json", err)
		return err
	}

	err = json.Unmarshal([]byte(file), &model.data)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return err
	}

	return nil
}

func (model *Model) LookUp(key string) (string, error) {
	ascii := model.data[key]
	if ascii == "" {
		return "", errors.New("Key not found")
	}
	return ascii, nil
}

func (model *Model) Add(key string, value string) (string, error) {
	/*
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(todo.Task + "\n")
		if err != nil {
			return err
		}
		return nil

	*/
}
