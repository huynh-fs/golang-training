package service

import (
	"fmt"
	"github.com/huynh-fs/interface-embedding/internal/domain"
)

type ZooService struct {
	creatures []domain.Creature
}

func NewZooService(creatures []domain.Creature) *ZooService {
	return &ZooService{creatures: make([]domain.Creature, 0)}
}

func (zs *ZooService) AddCreature(creature domain.Creature) {
	zs.creatures = append(zs.creatures, creature)
}

func (zs *ZooService) GenerateReport() {
	for _, creature := range zs.creatures {
		fmt.Printf("----Báo cáo về %s----\n", creature.GetName())
		fmt.Printf("Tiếng kêu: %s\n", creature.Speak())
		fmt.Printf("Cách di chuyển: %s\n", creature.Move())
		fmt.Println()
	}
}