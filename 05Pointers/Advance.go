package main

import "fmt"

type Player struct {
	health int
}

var healthy int 

func takeDamageFromExplosion(player Player) {
	fmt.Println("The player is about to take a damage, enter the intensity in numbers")
	var intensity int
	fmt.Scan(&intensity)
	player.health -= intensity
	fmt.Println("HEALTH in function is: ", player.health)
}

func main() {
	fmt.Println("Give a health to your player")
	fmt.Scan(&healthy)
	players := Player {
		health: healthy,
	}
	fmt.Println("Health of the BEFORE explosion player is: ", players.health)
	
	takeDamageFromExplosion(players)
	
	fmt.Println("Health of the AFTER explosion player is: ", players.health)
}
