package core

import "fmt"

type HandlerFunc func()

type StartupOptions struct {
	Id int
	Title string
	Handler HandlerFunc
}

func CreateNote () {
	fmt.Println("Create note")
}

func ShowAllNotes () {
	fmt.Println("Show all notes")
}

func Search () {
	fmt.Println("Search note")
}

func DeleteAllNotes () {
	fmt.Println("Delete all notes")
}

var StartupOptionsArray = []StartupOptions {
	StartupOptions{Id: 1, Title: "Create note", Handler: CreateNote, },
	StartupOptions{Id: 2, Title: "Show all notes", Handler: ShowAllNotes, },
	StartupOptions{Id: 3, Title: "Search for a note", Handler: Search, },
	StartupOptions{Id: 4, Title: "Delete all notes", Handler: DeleteAllNotes, },
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