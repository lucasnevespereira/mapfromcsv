package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type Data struct {
	arg1 string
	arg2 string
	arg3 string
	arg4 string
}

const (
	csvFileName = "example.csv"
	outFileName = "output.txt"
)

func main() {
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, col := range csvLines {
		d := Data{
			arg1: col[0],
			arg2: col[1],
			arg3: col[2],
			arg4: col[3],
		}
		writeOutput(d)
	}
}

func writeOutput(data Data) {
	f, err := os.OpenFile(outFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	funcName := strings.ToUpper(data.arg2)

	// func name string replacements
	for _, r := range Rules {
		funcName = strings.Replace(funcName, r.character, r.replacement, -1)
	}

	f.WriteString(
		funcName + "(" + "\"" +
			data.arg1 + "\"," +
			"\"" + data.arg2 + "\"," +
			"\"" + data.arg3 + "\"," +
			"\"" + data.arg4 +
			"\")" + ",\n")
}
