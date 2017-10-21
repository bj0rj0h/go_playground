package main

import (
	"fmt"
	"strings"
)

func main() {
	var i string
	var err error

	cont := true
	for cont {
		fmt.Print("Please enter a number\n")
		_, err = fmt.Scanf("%s", &i)
		if err == nil {
			switch i {
			case "1":
				one()
			case "2":
				fmt.Printf("What is your name?\n")
				var name string
				_, err = fmt.Scanf("%s", &name)
				two(name)
			case "3":
				three()
			case "4":
				four()
			case "5":
				primeTime()
			case "6":
				primeTime()
			default:
				fmt.Println("Not an option")
				cont = false
			}
		} else {
			fmt.Printf("Error occurred: %s\n", err)
		}
	}
}

func one() {
	fmt.Printf("Hello world!\n")
}

func two(name string) {
	if strings.EqualFold(name, "Alice") || strings.EqualFold(name, "Bob") {
		fmt.Printf("Hello %s\n", name)
	} else {
		one()
	}

}

func three() {
	//var num int
	fmt.Println("Please enter a number")

	var i int
	_, e := fmt.Scanf("%d", &i)
	if e != nil {
		err(e)
	} else {
		fmt.Printf("The sum is: %d\n", sumCalc(i))
	}
}

func err(e error) {
	fmt.Println("An error has occured")
	fmt.Println(e)
}

func sumCalc(num int) int {
	sum := 0
	for i := num; i > 0; i-- {
		sum += i
	}
	return sum
}

func sumCalcDiv36(num int) int {
	sum := 0
	for i := num; i > 0; i-- {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func getDigitFromUser() int {
	fmt.Println("Please enter a digit")
	i := 0
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error while getting int from user")
		fmt.Println(err)
	}
	return i
}

func four() {
	input := getDigitFromUser()
	printMultTable(input)
}

func printMultTable(num int) {
	for i := 0; i < 12; i++ {
		fmt.Printf("%d * %d = %d\n", i, num, num*i)
	}
}

func primeTime() {
	for i := 2; i < 1000; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(num int) bool {
	result := true
	if num > 2 {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				result = false
			}
		}
	} else {
		result = false
	}

	return result
}
