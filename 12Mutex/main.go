package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Player struct {
	health int
	mu sync.RWMutex
}

func newPlayer() *Player {
	
	return &Player{
		health: 100,
	}
}

func (players *Player) startUILoop() {
	// defer  wg.Done()
	ticker := time.NewTicker(time.Second)
	for {
		players.mu.RLock()
		fmt.Printf("Player health is: %v \n", players.health)
		<- ticker.C
		players.mu.RUnlock()
	}

}

func (players *Player) getDamage() {
	ticker := time.NewTicker(time.Millisecond * 300)
	i := 0
	rounds := []int{}
	for {
		
		players.mu.Lock()
		r:= rand.Intn(30)		
		players.health -= r

		println("R:", r)
		rounds = append(rounds, r)

		if players.health < 0 {
			fmt.Printf("THE PLAYER IS DEAD AFTER %v ROUNDS!!! \n", i)
			fmt.Println("The damage values are: ")
			for _, value := range rounds {
				fmt.Print(" ;", value)
			}			

			break
		}

		players.mu.Unlock()
		<- ticker.C
		i++
	}
}

func main() {
	player := newPlayer()	
	// wg.Add(1)
	go player.startUILoop()
	player.getDamage()
	// wg.Wait()
}

