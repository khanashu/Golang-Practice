package main

import (
	"fmt"

	"practice-go.com/bank-application/fileOps" //trying to use custom package i created

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println("Welcome to GO bank!")
	fmt.Println("Reach US 24X7", randomdata.PhoneNumber())
	var balance, err = fileOps.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println(err)
		fmt.Println("----------------")
		// panic(err) // Panic message with crash
	}
	var newBalance float64

	// for i := 0; i < 2; i++
	// for {
	// 	choice := optionSelect()
	// 	if choice == 1 {
	// 		balanceCheck(balance)
	// 	} else if choice == 2 {

	// 		newBalance = depositBalance(balance)
	// 		if newBalance < 0 {
	// 			fmt.Println("Invalid deposit amount Please try again!")
	// 		} else {
	// 			balance = newBalance
	// 			balanceCheck(balance)
	// 		}

	// 	} else if choice == 3 {
	// 		newBalance = withdrawBalance(balance)
	// 		if newBalance < 0 {
	// 			fmt.Println("Invalid withdrawn amount Please try again!")
	// 		} else {
	// 			balance = newBalance
	// 			balanceCheck(balance)
	// 		}
	// 	} else {
	// 		fmt.Println("Exiting Now!")
	// 		break
	// 		// return
	// 	}
	// }
	// fmt.Println("Thanks for choosing me")

	for {
		choice := optionSelect()
		switch choice {
		case 1:
			balanceCheck(balance)
		case 2:
			newBalance = depositBalance(balance)
			if newBalance < 0 {
				fmt.Println("Invalid deposit amount Please try again!")
				continue
			} else {
				balance = newBalance
				fileOps.WriteFloatToFile(balance, accountBalanceFile)
				balanceCheck(balance)
			}

		case 3:
			newBalance = withdrawBalance(balance)
			if newBalance < 0 {
				fmt.Println("Invalid withdrawn amount Please try again!")
				continue
			} else {
				balance = newBalance
				fileOps.WriteFloatToFile(balance, accountBalanceFile)
				balanceCheck(balance)
			}
		default:
			fmt.Println("Exiting Now!")
			fmt.Println("Thanks for choosing me")
			return
		}
	}

}

// func optionSelect() (choice int) {
// 	fmt.Println("What do you want to do?")
// 	fmt.Println("1. Check Balance")
// 	fmt.Println("2. Deposit Money")
// 	fmt.Println("3. Withdraw Money")
// 	fmt.Println("4. Exit")
// 	fmt.Print("Your Choice: ")
// 	fmt.Scan(&choice)
// 	fmt.Println("Selected menu:", choice)
// 	return choice
// }

func balanceCheck(balance float64) {
	fmt.Printf("Available Balance:%.1f\n", balance)
}

func depositBalance(balance float64) float64 {
	var deposit float64
	fmt.Println("Enter Amount to be Desposited: ")
	fmt.Scan(&deposit)
	if deposit < 0 {
		return -1
	}
	newBalance := balance + deposit
	return newBalance
}

func withdrawBalance(balance float64) float64 {
	var withdraw float64
	fmt.Println("Enter Amount to be Withdrawn: ")
	fmt.Scan(&withdraw)
	if withdraw > balance || withdraw < 0 {
		return -1
	}
	newBalance := balance - withdraw
	return newBalance
}
