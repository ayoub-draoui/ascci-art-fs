package main

import (
	"log"
	"os"
	"strings"

	ascciartfs "asciiartfs/functions"
)

func main() {
	readFile := map[rune][]string{}

	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalln("err: you shoud enter two or three argumrnts")
	}
	// if only two arguments
	if len(os.Args) == 2 {
		readFile = ascciartfs.ReadFile("./sources/standard.txt")
		// if it's three arguments
	} else if !strings.HasSuffix(os.Args[2], ".txt") {
		readFile = ascciartfs.ReadFile("./sources/" + os.Args[2] + ".txt")
		// if it doesn't contain .txt format
	} else if strings.HasSuffix(os.Args[2], ".txt") {
		readFile = ascciartfs.ReadFile("./sources/" + os.Args[2])
		// if it does contain .txt format
	} else {
		log.Fatalln("err: the third argument should be one of these file names (standerd),(shadow),(thinkertoy)")
	}
	// check if the input is among ascii manual

	checkcharacter := ascciartfs.CheckCharacter(os.Args[1])

	// go print my argument
	ascciartfs.FindAndPrint(checkcharacter, readFile)
}
