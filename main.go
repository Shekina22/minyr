package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			readcsv()
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til graderFahrenheit.")
		} else if input == "average" {
			fmt.Println("average")
		} else {
			fmt.Println("convert, average eller exit:")
		}

	}

}
