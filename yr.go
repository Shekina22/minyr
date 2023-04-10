package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readcsv(filename string) {
	src, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	log.Println(src)

	var fahr float64
	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
	bytesCount := 0
	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
			log.Println(string(linebuf))
			// Her
			elementArray := strings.Split(string(linebuf), ";")
			if len(elementArray) > 3 {
				celsius := elementArray[3]
				fahr := celsiusToFahrenheit(float64(celsius))
				log.Println(elementArray[3])
			}
			linebuf = nil
		} else {
			linebuf = append(linebuf, buffer[0])
		}
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}

	// Loop through the list of temperatures and convert each one to Fahrenheit
	for _, temp := range temperatures {
		celsius := temp[0]
		fahrenheit := celsiusToFahrenheit(celsius)
		fmt.Printf("%.0f°C = %.1f°F\n", celsius, fahrenheit)
	}
}

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius - 32) * 5 / 9
}
