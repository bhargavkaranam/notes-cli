package core

import (
	"fmt"
	"github.com/bhargavkaranam/notes-cli/fs"
	"path/filepath"
	"os/exec"
)

const DIRECTORY_PATH = "Notes/"

type HandlerFunc func()

type StartupOptions struct {
	Id int
	Title string
	Handler HandlerFunc
}

func createNote () {
	var title string

	fmt.Println("Enter a title to easily find the note")
	fmt.Scan(&title)

	constructedFilePath, errorInConstructionFilePath := filepath.Abs(DIRECTORY_PATH + title + ".md")

	if (errorInConstructionFilePath != nil) {
		panic(errorInConstructionFilePath)
	}

	filename, err := fs.CreateFile(constructedFilePath)

	if err != nil {
		panic(err)
	}

	fmt.Println("Add the note in your favorite text editor save it!", filename, constructedFilePath)

	exec.Command("code", constructedFilePath).Run()
}

func showAllNotes () {
	fmt.Println("Show all notes")
}

func search () {
	fmt.Println("Search note")
}

func deleteAllNotes () {
	fmt.Println("Delete all notes")
}

var StartupOptionsArray = []StartupOptions {
	StartupOptions{Id: 1, Title: "Create note", Handler: createNote, },
	StartupOptions{Id: 2, Title: "Show all notes", Handler: showAllNotes, },
	StartupOptions{Id: 3, Title: "Search for a note", Handler: search, },
	StartupOptions{Id: 4, Title: "Delete all notes", Handler: deleteAllNotes, },
}

func GetStartupOptions () []StartupOptions {
	return StartupOptionsArray
}

func PerformStartupAction (Id int) {
	for _, option := range StartupOptionsArray {
		if option.Id == Id {
			option.Handler()
			break
		}
	}
}