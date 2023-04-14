package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Shekina22/minyr/yr"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = strings.ToLower(scanner.Text())
		if input == "exit" {
			os.Exit(0)

		} else if input == "convert" {
			if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
				fmt.Print("Generere på nytt? (y/n): ")
				scanner.Scan()
				answer := strings.ToLower(scanner.Text())
				if answer == "y" {
					convertedTemperatures, err := yr.ConvertTemperatures()
					if err != nil {
						log.Fatal(err)
					}

					if err := yr.SkrivLinjer(convertedTemperatures, "kjevik-temp-fahr-20220318-20230318.csv"); err != nil {
						log.Fatal(err)
					}

					fmt.Println("Filen er konvert.")
				} else if answer == "n" {
					return
				} else {
					log.Fatal("Velg Y eller N")
					return
				}
			}

			convertedTemperatures, err := yr.ConvertTemperatures()
			if err != nil {
				log.Fatal(err)
			}

			if err := yr.SkrivLinjer(convertedTemperatures, "kjevik-temp-fahr-20220318-20230318.csv"); err != nil {
				log.Fatal(err)
			}

		} else if input == "average" {
			var unit string
			for unit != "c" && unit != "f" {
				fmt.Print("Hvil du ha utregningen i Celsius eller Fahrenheir? (c / f): ")
				scanner.Scan()
				unit = strings.ToLower(scanner.Text())
			}

			file, err := yr.ÅpneFil("kjevik-temp-celsius-20220318-20230318.csv")
			if err != nil {
				log.Fatal(err)
			}
			defer yr.LukkFil(file)

			lines, err := yr.LesLinjer(file)
			if err != nil {
				log.Fatal(err)
			}

			var sum float64
			count := 0
			for i, line := range lines {
				if i == 0 {
					continue
				}
				fields := strings.Split(line, ";")
				if len(fields) != 4 {
					log.Fatalf("Feil med felt i linje %d: %d", i, len(fields))
				}
				if fields[3] == "" {
					continue
				}
				if temperature, err := strconv.ParseFloat(fields[3], 64); err != nil {
					log.Fatalf("Feil med temperatur i linje %d: %s", i, err)
				} else {
					if unit == "f" {
						temperature = temperature*1.8 + 32
					}
					sum += temperature
					count++
				}
			}
			average := sum / float64(count)

			if unit == "f" {
				fmt.Printf("Gjennomsnittlig temperatur i fahrenheit: %.2f°F\n", average)
			} else {
				fmt.Printf("Gjennomsnittlig temperatur i celsius: %.2f°C\n", average)
			}
		} else {
			fmt.Println("CONVERT, AVERAGE eller EXIT")
		}
	}

}
