package main

import "fmt"

func main() {
	var i string = ""
	var err error

	cont := true

	for cont {
		_, err = fmt.Scanf("%s", &i)
		if err == nil {
			switch i {
			case "1":
				fmt.Println("one")
			case "2":
				fmt.Println("two")
			default:
				fmt.Println("Not an option")
				cont = false
			}
		} else {
			fmt.Printf("Error occurred: %s\n", err)
		}
	}
}
