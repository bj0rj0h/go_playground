package main

import "fmt"

func main() {
	var i int = 0
	_, err := fmt.Scanf("%d", &i)

	if err == nil {
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println("Not an option")
		}
	} else {
		fmt.Print("Error occured: ")
		fmt.Println(err)
	}

	fmt.Scanf("%d", i)
}
