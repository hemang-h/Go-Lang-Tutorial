package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Player struct {
	health int
}

func newPlayer() *Player {
	
	return &Player{
		health: 100,
	}
}

func (players *Player) startUILoop() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for {
		players.health -= rand.Intn(40)
		if players.health < 0 {
			fmt.Printf("THE PLAYER IS DEAD AFTER %v ROUNDS!!! \n", i)
			break
		}
		fmt.Printf("Player health is: %d \r", players.health)
		<- ticker.C
		i++
	}
}

func main() {
	// players := &Player {
	// 	health: ,
	// }
	newPlayer().startUILoop()
	

}

