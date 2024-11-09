package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"go-practice.com/noteApp/note"
)

func main() {

	noteTitle, noteContent, err := getNoteData()
	if err != nil {
		return
	}

	var anyNote note.Note
	anyNote, err = note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = anyNote.Save()
	if err != nil {
		fmt.Println("Saving failed!")
		return
	}

	fmt.Println("Saving Success!")
}

func getNoteData() (string, string, error) {
	noteTitle, err := getUserData("Title of Note:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	noteContent, err := getUserData("Note Content:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return noteTitle, noteContent, nil
	// var anyNote *note.Note
	// anyNote, err := note.AddNote(noteTitle, noteContent)
}

func getUserData(message string) (string, error) {
	fmt.Printf("%v ", message)
	// fmt.Scanln(&value) can't deal with spaces
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New(" Title and Content can't be empty")
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text, nil
}
