package function

import (
	"fmt"
	"strconv"
	"strings"
)

//ParseInput is method for get data by element
func ParseInput(input string) (data []int, err bool) {
	temp := strings.Replace(input, "[", "", -1)
	temp = strings.Replace(temp, "]", "", -1)
	temp = strings.Replace(temp, " ", "", -1)
	temp = strings.Replace(temp, "\n", "", -1)
	results := strings.Split(temp, ",")
	data, err = convert2Int(results)
	return data, err
}

func convert2Int(input []string) (output []int, e bool) {
	for _, i := range input {
		j, err := strconv.Atoi(i)
		if err != nil {
			e = true
			return output, e
		} else {
			e = false
			output = append(output, j)
		}
	}
	return output, e
}

func visualize(input []int) {
	max := maxValue(input)

	for j := max; j >= 0; j-- {
		for _, i := range input {
			if i > j {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}

	for _, i := range input {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func maxValue(input []int) (output int) {
	max := input[0]
	for _, i := range input {
		if max < i {
			max = i
		}
	}
	return max
}

// InsertionSort is method for sort array input integer
func InsertionSort(inputA []int, n int, sort string) (outputA []int) {
	if n > 0 {
		InsertionSort(inputA, n-1, sort)
		x := inputA[n]
		j := n - 1
		if sort == "A" {
			for j >= 0 && inputA[j] > x {
				inputA[j+1] = inputA[j]
				j = j - 1
			}
		} else if sort == "D" {
			for j >= 0 && inputA[j] < x {
				inputA[j+1] = inputA[j]
				j = j - 1
			}
		}
		inputA[j+1] = x
	}
	if n == 0 {
		fmt.Println("-----ORIGINAL DATA : INSERTION SORT : ACCENDING------")
	} else {
		fmt.Println("-----STEP", n, "------")
	}
	visualize(inputA)
	return inputA
}
