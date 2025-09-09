package random

import (
	"math/rand"
	"fmt"
)

type Generator struct {
	draw map[int]bool
}

func NewGenerator() *Generator {
	return &Generator{
		draw: make(map[int]bool),
	}
}

func (g *Generator) Draw(min, max int) (int, error) {
	if len(g.draw) >= (max - min + 1) {
		return 0, fmt.Errorf("tất cả số đã được rút")
	}
	for {
		num := rand.Intn(max-min+1) + min
		if !g.draw[num] {
			g.draw[num] = true
			return num, nil
		}
	}
}