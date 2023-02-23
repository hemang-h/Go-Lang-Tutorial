package main

import "fmt"

type Player struct {
	health int
	Name string
}


var healthy int 

func (player Player) takeDamageFromExplosion(intensity int) {                                 //We can declare the function like this
	fmt.Println("Function before the number hit and enter a intensity")
	player.health -= (intensity + 5)
	fmt.Println("HEALTH in function is: ", player.health)
}

func takeDamageFromExplosion(player *Player) {                                               // player *Player shows that I am referencing the value in the memory 
	fmt.Println("The player is about to take a damage, enter the intensity in numbers")
	var intensity int
	fmt.Scan(&intensity)
	player.health -= intensity
	fmt.Println("HEALTH in function is: ", player.health)
	fmt.Println("NAME in function is: ", player.Name)
}

func main() {
	fmt.Println("Give a health to your player")
	fmt.Scan(&healthy)
	players := &Player {                                                                       //Here I have initialized the player directly as a pointer
		health: healthy,
		Name: "Hemang",
	}
	fmt.Println("Health of the BEFORE explosion player is: ", players.health)
	
	players.takeDamageFromExplosion(10)
	takeDamageFromExplosion(players)                                                         // I can also pass the same value as takeDamageFromExplosion(&players) note the & here and compare with parameter passed in the function 
	
	fmt.Println("Health of the AFTER explosion player is: ", players.health)
}
