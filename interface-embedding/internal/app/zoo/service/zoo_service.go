package service

import (
	"fmt"
	"github.com/huynh-fs/interface-embedding/pkg/animal"
)

type ZooService struct {
	creatures []animal.Creature
}

func NewZooService() *ZooService {
	return &ZooService{creatures: make([]animal.Creature, 0)}
}

func (s *ZooService) AddCreature(c animal.Creature) {
	s.creatures = append(s.creatures, c)
}

func (s *ZooService) GenerateReport() {
	fmt.Println("==============================")
	for _, c := range s.creatures {
		fmt.Printf("--- Báo cáo về %s ---\n", c.GetName())
		fmt.Printf("Tiếng kêu: %s\n", c.Speak())
		fmt.Printf("Cách di chuyển: %s\n", c.Move())
		fmt.Println()
	}
	fmt.Println("==============================")
}