default: help

help:
	@echo "select a subcommand:\n"
	@echo "run:\t run webserver"
	@echo "build:\t build webserver"

run:
	go run main.go

build:
	go build -o gsearch-web