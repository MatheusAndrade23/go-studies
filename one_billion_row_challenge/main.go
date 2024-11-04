package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	measurements, err := os.Open("measurements.txt")

	if err != nil {
		panic(err)
	}

	defer measurements.Close()

	dados := make(map[string]Measurement)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		line := scanner.Text()
		semicolon := strings.Index(line, ";")
		location := line[:semicolon]
		rawTemp := line[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)
		
		measurement, ok := dados[location]

		if !ok {
			measurement = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measurement.Min = min(measurement.Min, temp)
			measurement.Max = max(measurement.Max, temp)
			measurement.Sum += temp
			measurement.Count++
		}
		dados[location] = measurement
	}

	locations := make([]string, 0, len(dados))

	for location := range dados {
		locations = append(locations, location)
	}

	sort.Strings(locations)

	fmt.Printf("{")
	for _, location := range locations {
		measurement := dados[location]
		avg := measurement.Sum / float64(measurement.Count)
		fmt.Printf("%s=%.1f/%.1f/%.1f, ", location, measurement.Min, avg, measurement.Max)
	}
	fmt.Printf("}")
}