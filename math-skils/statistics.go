package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"math"
)

func StandardDeviation(s string) int {
	return int(math.Sqrt(float64(Variance(s))))
}

func sortIntegers(numbers []int) {

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

func AverageFLoat(s string) float64 {

	numbersString := strings.Split(s, "\n")

	var numbers []int

	for _, element := range numbersString {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, n)
	}

	somme := 0.0
	for _, number := range numbers {
		somme += float64(number)
	}

	return float64(somme/float64(len(numbers)))
}

func Variance(s string) int {

	numbersString := strings.Split(s, "\n")

	var numbers []int

	for _, element := range numbersString {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, n)
	}

	var squareNumbers []float64
	for _, number := range numbers {

		squareNumbers = append(squareNumbers, math.Pow( (float64(number) - float64(AverageFLoat(s))), float64(2)))
	}

	sommeSquareNumber := 0.0

	for _, number := range squareNumbers {
		sommeSquareNumber += number
	}
	return int(sommeSquareNumber / float64(len(numbers)))
}

func Average(s string) int {

	numbersString := strings.Split(s, "\n")

	var numbers []int

	for _, element := range numbersString {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, n)
	}

	somme := 0
	for _, number := range numbers {
		somme += number
	}

	return int(somme/len(numbers))
}

func Median(s string) int {
	
	numbersString := strings.Split(s, "\n")

	var numbers []int

	for _, element := range numbersString {
		n, _ := strconv.Atoi(element)
		numbers = append(numbers, n)
	}

	sortIntegers(numbers)

	median := 0

	if len(numbers) % 2 == 0 {
		median = int(numbers[(len(numbers) / 2) - 1])
	} else {
		median = int(numbers[(len(numbers) / 2)])
	}

	return median
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

	average := Average(sampleFileString)
	median := Median(sampleFileString)
	variance := Variance(sampleFileString)
	std := StandardDeviation(sampleFileString)
	fmt.Println(std)
	floatAverage := AverageFLoat(sampleFileString)
	fmt.Println(floatAverage)
	fmt.Println(variance)
	fmt.Println(median)
	fmt.Println(average)
}