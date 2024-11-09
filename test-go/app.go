package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}
func main() {
	fmt.Print("Hello World")
	var name customString = "Ashu"

	name.log()

}
