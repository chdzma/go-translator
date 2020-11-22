package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bregydoc/gtranslate"
)

const fromLanguage string = `en`
const toLanguage string = `es`
const csvFileForRead string = `example.csv`
const csvFileForWrite string = `translate.csv`

func main() {
	records := translateCsvFile()
	writeCsvFile(records)
}

func translateCsvFile() [][]string {
	records := [][]string{}
	csvfile, err := os.Open(csvFileForRead)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		original := fmt.Sprintf(`%s`, record[0])
		translate := translate(record[1])
		fmt.Println(fmt.Sprintf(`%s --> %s`, original, translate))
		records = append(records, []string{original, translate})
	}
	return records
}

func translate(text string) string {
	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: fromLanguage,
			To:   toLanguage,
		},
	)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf(`%s`, translated)
}

func writeCsvFile(records [][]string) {
	f, err := os.Create(csvFileForWrite)
	defer f.Close()

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(f)
	err = w.WriteAll(records)

	if err != nil {
		log.Fatal(err)
	}
}
