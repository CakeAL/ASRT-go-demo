default:
    @just --list

alias r := run
alias b := build

run:
    go mod tidy
    go run main.go -f test.wav

build: 
    go mod tidy
    go build