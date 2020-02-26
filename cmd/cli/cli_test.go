package main

import (
	"testing"

	poker "go.learning.application/poker"
)

func CLITest(t *testing.T) {
	playerStore := &poker.StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	cli := &CLI{playerStore}
	cli.PlayPoker()

	if len(playerStore.winCalls) != 1 {
		t.Fatal("expected a win but didn't get any")
	}
}
