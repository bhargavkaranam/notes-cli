module github.com/bhargavkaranam/notes-cli/consumer

go 1.20

replace github.com/bhargavkaranam/notes-cli/core => ./core

require github.com/bhargavkaranam/notes-cli/core v0.0.0-00010101000000-000000000000

require github.com/bhargavkaranam/notes-cli/fs v0.0.0-00010101000000-000000000000 // indirect

replace github.com/bhargavkaranam/notes-cli/fs => ./fs
