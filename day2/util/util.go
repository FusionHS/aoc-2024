package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadReports() ([][]int, error) {
	//file, err := os.Open("sample-a.txt")
	file, err := os.Open("input-a.txt")
	if err != nil {
		log.Fatalf("failed to open the file: %s", err)
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	rows := strings.Split(string(content), "\n")

	var reports [][]int

	for _, row := range rows {
		columns := strings.Fields(row)

		var intValues []int
		for _, column := range columns {
			num, err := strconv.Atoi(column)

			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil, err
			}
			intValues = append(intValues, num)
		}

		reports = append(reports, intValues)
	}

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return reports, nil
}
