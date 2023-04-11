package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func GetAverageLufttemperatur(filename string) (float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Skip the header line
	_, err = reader.Read()
	if err != nil {
		return 0, err
	}

	sum := 0
	count := 0

	// Read each record and calculate the sum of Lufttemperatur values
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		lufttemperatur, err := strconv.Atoi(record[3]) // Assuming Lufttemperatur is at index 3
		if err != nil {
			return 0, err
		}

		sum += lufttemperatur
		count++
	}

	average := float64(sum) / float64(count)

	return average, nil
}

func ConvertCelsiusToFahrenheit(filename string, createNewFile bool) error {
	// Open the original CSV file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the original CSV data
	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	var newFilename string
	if createNewFile {
		// Create a new CSV file for writing
		newFilename = "converted_" + filename
		newFile, err := os.Create(newFilename)
		if err != nil {
			return err
		}
		defer newFile.Close()

		// Create a CSV writer
		writer := csv.NewWriter(newFile)
		writer.Comma = ';'

		// Write the header line
		header := []string{"Navn", "Stasjon", "Tid(norsk normaltid)", "Lufttemperatur(Fahrenheit)"}
		err = writer.Write(header)
		if err != nil {
			return err
		}

		// Convert and write each record
		for _, record := range records {
			lufttemperatur, err := strconv.Atoi(record[3]) // Assuming Lufttemperatur is at index 3
			if err != nil {
				return err
			}

			fahrenheit := CelsiusToFahrenheit(float64(lufttemperatur))
			record = append(record, strconv.FormatFloat(fahrenheit, 'f', 2, 64)) // Append the converted value
			err = writer.Write(record)
			if err != nil {
				return err
			}
		}

		writer.Flush()

		fmt.Println("Conversion complete. Converted data written to:", newFilename)
	} else {
		// Overwrite the original file with the converted data
		newFilename = filename
		newFile, err := os.Create(newFilename)
		if err != nil {
			return err
		}
		defer newFile.Close()

		// Create a CSV writer
		writer := csv.NewWriter(newFile)
		writer.Comma = ';'

		// Write the header line
		header := []string{"Navn", "Stasjon", "Tid(norsk normaltid)", "Lufttemperatur(Fahrenheit)"}
		err = writer.Write(header)
		if err != nil {
			return err
		}

		// Convert and write each record
		for _, record := range records {
			lufttemperatur, err := strconv.Atoi(record[3]) // Assuming Lufttemperatur is at index 3
			if err != nil {
				return err
			}

			fahrenheit := CelsiusToFahrenheit(float64(lufttemperatur))
			record[3] = strconv.FormatFloat(fahrenheit, 'f', 2, 64) // Update the converted value
			err = writer.Write(record)
			if err != nil {
				return err
			}
		}

		writer.Flush()

		fmt.Println("Conversion complete. Overwritten data written to:", newFilename)
	}

	return nil
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius - 32) * 5 / 9
}
