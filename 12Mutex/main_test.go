package main

import "testing"

func TestGame(t *testing.T) {
	player := newPlayer()	
	go player.startUILoop()
	player.getDamage()	
}	