//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import (
	"fmt"
	"testing"
)

func TestPlayerHealthLtMax(t *testing.T) {
	player := newPlayer("Nils", 5, 5)
	player.addHealth(1)
	if player.health != 5 {
		t.Errorf(fmt.Sprintln(
			"Health cannot be above max health. Was",
			player.health,
			"when max was",
			player.maxHealth,
		),
		)
	}
}
func TestPlayerHealthGtZero(t *testing.T) {
	player := newPlayer("Nils", 5, 5)
	player.addHealth(-10)
	if player.health != 0 {
		t.Errorf(fmt.Sprintln(
			"Health cannot be negative. Was",
			player.health,
			"when min was 0",
		),
		)
	}
}
