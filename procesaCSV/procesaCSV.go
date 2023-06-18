package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func openCSV(csvname string) (io.Reader, error) {
	file, err := os.Open(csvname)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func scan(csvreader io.Reader) func() ([]interface{}, error) {
	scanner := bufio.NewScanner(csvreader)
	return func() ([]interface{}, error) {
		for scanner.Scan() {
			iline := scanner.Text()
			dataline := make([]interface{}, 0)
			//fields := reSep.Split(iline, -1)
			fields := regexp.MustCompile(";").Split(iline, -1)
			for _, field := range fields {
				/*switch {
				case reInteger.MatchString(field):
					value, _ := strconv.Atoi(field)
					dataline = append(dataline, value)
				case reFloat.MatchString(field):
					value, _ := strconv.ParseFloat(field, 64)
					dataline = append(dataline, value)
				default:
					dataline = append(dataline, field)
				}*/
				dataline = append(dataline, field)
			}
			return dataline, nil
		}
		return nil, errors.New("EOF")
	}
}

func main() {
	if csvfile, err := openCSV("music.csv"); err != nil {
		log.Fatal(err)
	} else {
		scanner := scan(csvfile)
		for data, err := scanner(); err == nil; data, err = scanner() {
			for _, field := range data {
				switch field.(type) {
				case int:
					fmt.Printf("%d (int)", field)
				case float64:
					fmt.Printf("%v (float)", field)
				case string:
					fmt.Printf("%v (string)", field)
				}
			}
			fmt.Println()
		}
	}
}
