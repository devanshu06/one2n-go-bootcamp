package main

import (
	"fmt"
)

var nums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func evenNumbers(number []int) []int {
	var output []int
	for _, n := range number {
		if n%2 == 0 {
			output = append(output, n)
		}
	}
	return output
}

func oddNumbers(number []int) []int {
	var output []int
	for _, n := range number {
		if n%2 != 0 {
			output = append(output, n)
		}
	}
	return output
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeNumbers(number []int) []int {
	var output []int
	for _, n := range number {
		if isPrime(n) {
			output = append(output, n)
		}
	}
	return output
}

func oddPrimeNumbers(number []int) []int {
	var output []int
	for _, n := range number {
		if isPrime(n) && n%2 != 0 {
			output = append(output, n)
		}
	}
	return output
}

func evenMutlipleoffive(number []int) []int {
	var output []int
	for _, n := range number {
		if n%2 == 0 && n%5 == 0 {
			output = append(output, n)
		}
	}
	return output
}

func oddMutlipleofthree(number []int) []int {
	var output []int
	for _, n := range number {
		if n%6 == 3 && n > 10 {
			output = append(output, n)
		}
	}
	return output
}

func main() {
	fmt.Println("Hello world!!")
	fmt.Println("Even Numbers: ", evenNumbers(nums))
	fmt.Println("Odd Numbers: ", oddNumbers(nums))
	fmt.Println("Prime Numbers: ", primeNumbers(nums))
	fmt.Println("Odd Prime Numbers:", oddPrimeNumbers(nums))
	fmt.Println("Even Number with the multiple of 5: ", evenMutlipleoffive(nums))
	fmt.Println("Odd Number with the multiple of 3 and Value greater than 10: ", oddMutlipleofthree(nums))

}
