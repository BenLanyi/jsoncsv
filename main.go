package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		panic("Must provide path")
	}

	path := args[0]

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	valuesMap := map[string]string{}
	err = json.Unmarshal(bytes, &valuesMap)
	if err != nil {
		panic(err)
	}

	outputFile, err := os.Create("translate.csv")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	writer.Comma = '|'

	record := []string{"text", "translation"}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}

	for key := range valuesMap {
		fmt.Println(key)
		record := []string{fmt.Sprintf("%s", key), ""}
		err = writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

}
