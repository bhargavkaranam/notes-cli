package main

import (
	"fmt"
	"github.com/bhargavkaranam/notes-cli/core"
	"github.com/bhargavkaranam/notes-cli/fs"
)

func bootstrap () {
	fs.CreateDirectoryIfNotExists("Notes")
}

func main () {
	bootstrap()
	fmt.Println("Welcome to the tiny note taking application.")

	var all_options = core.GetStartupOptions()
	
	for _, option := range all_options {
		fmt.Printf("%d. %v\n", option.Id, option.Title)
	}

	var user_option int

	fmt.Print("Enter an option:")
	fmt.Scan(&user_option)

	core.PerformStartupAction(user_option)
}