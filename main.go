package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	switch {
	case strings.HasSuffix(file.Name(), ".csv"):
		err := csvToJson(file)
		if err != nil {
			panic(err)
		}

	case strings.HasSuffix(file.Name(), ".json"):
		err := jsonToCsv(file)
		if err != nil {
			panic(err)
		}
	default:
		panic("Unsupported file type. Must be CSV or JSON")
	}

}

func jsonToCsv(file *os.File) error {
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	valuesMap := map[string]string{}
	err = json.Unmarshal(bytes, &valuesMap)
	if err != nil {
		return err
	}

	outputFile, err := os.Create("translate-this.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	writer.Comma = '|'

	record := []string{"text", "translation"}
	err = writer.Write(record)
	if err != nil {
		return err
	}

	for key := range valuesMap {
		record := []string{fmt.Sprintf("%s", key), ""}
		err = writer.Write(record)
		if err != nil {
			return err
		}
	}

	fmt.Println("output to", outputFile.Name())

	return nil
}

func csvToJson(file *os.File) error {
	reader := csv.NewReader(file)
	reader.Comma = '|'

	values, err := reader.ReadAll()
	if err != nil {
		return err
	}

	outputFile, err := os.Create("translated.json")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString("{\n")
	if err != nil {
		return err
	}

	for key, value := range values {
		if value[0] == "text" && value[1] == "translation" {
			continue
		}

		comma := ""
		if key != len(values)-1 {
			comma = ","
		}

		_, err := outputFile.WriteString(fmt.Sprintf("  \"%s\": \"%s\"%s\n", value[0], value[1], comma))
		if err != nil {
			return err
		}
	}

	_, err = outputFile.WriteString("}")
	if err != nil {
		return err
	}

	fmt.Println("output to", outputFile.Name())

	return nil
}
