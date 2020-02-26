package main

import (
	"fmt"

	poker "go.learning.application/poker"
)

type CLI struct {
	playerStore poker.PlayerStore
}

func (cli *CLI) PlayPoker() {
	cli.playerStore.RecordWin("Cleo")
}

func main() {
	fmt.Println("Let's play poker")
}
