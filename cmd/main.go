package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func commandCleaning(c string) string {

	return strings.ToLower(strings.TrimSpace(c))

}

func main() {
	file := flag.String("file", "", "name of file or full path for command")
	command := flag.String("do", "", "command for file treatment: create, read, delete")

	flag.Parse()

	var emptyValue string

	if *file == emptyValue {
		log.Fatal("Do not exist file name or file path. Tool is clousing")
	}

	if *command == emptyValue {
		log.Fatal("Command is absent. Try one more time from the begining. Application is shooting down.")
	}

	cleanedCommand := commandCleaning(*command)

	switch cleanedCommand {

	case emptyValue:
		log.Fatal("Command is absent. Try one more time from the begining. Application is shooting down.")

	case "create":
		create(*file)

	case "read":
		read(*file)

	case "delete":
		remove(*file)

	default:
		fmt.Printf("Unknow command %s. Start application with flag -help. Read manual", cleanedCommand)
	}

}
