package main 

type Player struct {
	health int
}

func newPlayer() *Player {
	return &Player{
		health: 100,
	}
} 
