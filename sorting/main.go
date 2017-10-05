package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rezandry/prescreening/sorting/function"
)

// Main is main function from this program
func main() {
	// Input data
	fmt.Println("Input your array number, split with comma : ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Parse input data
	data, err := function.ParseInput(text)

	// Check if input data is not integer and do insertion sort if integer
	if err == true {
		fmt.Println("Your input must be integer!")
	} else {
		var x int
		fmt.Println("Sort data by : ")
		fmt.Println("1. Accending", "\n2. Deccending")
		fmt.Scan(&x)
		if x == 1 {
			dataSort := function.InsertionSort(data, len(data)-1, "A")
			fmt.Println("Insertion Sort : ", dataSort)
		} else if x == 2 {
			dataSort := function.InsertionSort(data, len(data)-1, "D")
			fmt.Println("Insertion Sort : ", dataSort)
		} else {
			fmt.Println("Your input must be 1 or 2!")
		}
	}
}
