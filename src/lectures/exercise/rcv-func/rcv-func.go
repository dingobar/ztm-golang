//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Player struct {
	name      string
	health    int
	maxHealth int
	energy    int
	maxEnergy int
}

func (player *Player) addHealth(diff int) {
	player.health += diff
	if player.health > player.maxHealth {
		player.health = player.maxHealth
	} else if player.health < 0 {
		player.health = 0
	}
}
func (player *Player) addEnergy(diff int) {
	player.energy += diff
	if player.energy > player.maxEnergy {
		player.energy = player.maxEnergy
	} else if player.energy < 0 {
		player.energy = 0
	}
}

func (player *Player) attack(other *Player) {
	damage := rand.Intn(player.energy)
	fmt.Println(player.name, "attacks", other.name, "for", damage, "damage")
	other.addHealth(-damage)
}

func newPlayer(name string, maxHealth int, maxEnergy int) Player {
	return Player{
		name:      name,
		health:    maxHealth,
		maxHealth: maxHealth,
		energy:    maxEnergy,
		maxEnergy: maxEnergy,
	}
}

func main() {
	player1 := newPlayer("Joe", 100, 10)
	player2 := newPlayer("Morty", 50, 20)

	for i := 0; i < 100; i++ {
		player1.attack(&player2)
		if player2.health == 0 {
			break
		}
		player2.attack(&player1)
		if player1.health == 0 {
			break
		}
	}

	fmt.Println("The fight is over.", player1, player2)
}
