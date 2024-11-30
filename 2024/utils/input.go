package utils

import (
	"bufio"
	"log"
	"os"
)

func GetData(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
