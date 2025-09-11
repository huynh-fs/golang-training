package service

import (
	"fmt"
	"math/rand"
)

type RandomService struct {
	drawn map[int]bool
}


func NewRandomService() *RandomService {
	return &RandomService{
		drawn: make(map[int]bool),
	}
}

func (s *RandomService) Draw(min, max int) (int, error) {
	if len(s.drawn) >= (max - min + 1) {
		return 0, fmt.Errorf("tất cả các số trong khoảng [%d, %d] đã được rút", min, max)
	}

	for {
		num := rand.Intn(max-min+1) + min
		if !s.drawn[num] {
			s.drawn[num] = true
			return num, nil
		}
	}
}