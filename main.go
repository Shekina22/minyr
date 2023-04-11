package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("convert, average eller exit:")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	input = scanner.Text()
	for scanner.Scan() {
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Do you want to create a new file?")
			if input == "Y" {
				ConvertCelsiusToFahrenheit("kjevik-temp-celsius-20220318-20230318.csv", true)
			} else if input == "n" {
				ConvertCelsiusToFahrenheit("kjevik-temp-celsius-20220318-20230318.csv", false)
			}
		} else if input == "average" {
			if input == "f" {
				filename := "kjevik-temp-celsius-20220318-20230318.csv"
				average, err := GetAverageLufttemperatur(filename)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Println("Lufttemperatur", average)
			} else if input == "c" {
				filename := "kjevik-temp-celsius-20220318-20230318.csv"
				average, err := GetAverageLufttemperatur(filename)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}
				fmt.Println("Lufttemperatur", average)
			}
		} else {
			fmt.Println("convert, average eller exit:")
		}

	}
}
