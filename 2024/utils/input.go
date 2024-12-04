package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func MapToInt(s []string) []int {
	intSlice := make([]int, 0, len(s))
	for _, el := range s {
		intEl, _ := strconv.Atoi(el)
		intSlice = append(intSlice, intEl)
	}
	return intSlice
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
