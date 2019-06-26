package main

import (
	"fmt"
	"github.com/friendsofgo/board-kata"
	"github.com/friendsofgo/board-kata/parser"
	"log"
	"os"
)

func main() {
	msg, err := board.ReadInput("data/input.csv")

	if board.IsImpossibleOpenFile(err) {
		log.Fatal(err)
	}

	file, err := os.Create("data/result.html")
	defer file.Close()

	for _, currentMsg := range msg {
		output := parser.Parse(currentMsg)
		fmt.Println(output)
		file.WriteString(output)
	}

	fmt.Println("done!")
}
