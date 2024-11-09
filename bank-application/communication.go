package main

import "fmt"

func optionSelect() (choice int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Your Choice: ")
	fmt.Scan(&choice)
	fmt.Println("Selected menu:", choice)
	return choice
}
