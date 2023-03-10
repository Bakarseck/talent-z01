package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func Average(tabFloat []float64) float64 {

	sum := 0.0
	for _, v := range tabFloat {
		sum += v
	}
	return sum / float64(len(tabFloat))
}

func Median(tabFloat []float64) float64 {

	sortFloat64(tabFloat)

	length := len(tabFloat)

	middle := len(tabFloat) / 2

	median := 0.0

	if length % 2 == 0 {
		middleLeft := tabFloat[middle -1]
		middleRight := tabFloat[middle]
		median = (middleLeft + middleRight) / 2
	} else if length % 2 != 0{
		median = tabFloat[middle]
	}
	return median
}

func Variance(tabFloat []float64, mean float64) float64 {

	sum := 0.0

	for _, v := range tabFloat {
		sum += math.Pow(v - mean, 2)
	}

	return sum / float64(len(tabFloat))
}

func StandardDeviation(tabFloat []float64, mean float64) float64 {
 	return math.Sqrt(float64(Variance(tabFloat, mean)))
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	arguments := os.Args[1:]

	sampleFile, err := ioutil.ReadFile(arguments[0])

	check(err)

	sampleFileString := string(sampleFile)

	tabFloat := stringToFloat(sampleFileString)

	average := Average(tabFloat)
	median := Median(tabFloat)
	variance := Variance(tabFloat, average)
	std := StandardDeviation(tabFloat, average)
	fmt.Printf("Mean: %v\n", int(math.Round(average)))
	fmt.Printf("Median: %v\n", int(math.Round(median)))
	fmt.Printf("Variance: %v\n", int(math.Round(variance)))
	fmt.Printf("Standard deviation: %v\n", int(math.Round(std)))
}

func sortFloat64(numbers []float64) {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
}

func stringToFloat(str string) []float64 {

	s := strings.TrimSpace(str)

	numbersString := strings.Split(s, "\n")

	var numbers []float64

	for _, element := range numbersString {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, float64(n))
	}
	return numbers
}
