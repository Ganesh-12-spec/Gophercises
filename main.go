package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

var filename = flag.String("file", "problems.csv", "path to csv file containing quiz problems")

func main() {
	file , err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records,err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading csv:",err)
		return
	}
	correct := 0
	for _, record := range records {
		question := record[0]
		answer := record[1]

		fmt.Print(question, "=")

		var userAnswer string
		fmt.Scanln(&userAnswer)

		if userAnswer == answer {
			correct++
		}
		}
		fmt.Println("you got" , correct, "out of", len(records))
	}