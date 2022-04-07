package model

import (
	"sync"

	"github.com/vokestd/TestingSystem-server/stuff"
)

var traitsOnce sync.Once

type Trait struct {
	Name         string
	FriendlyName string
	Description  string
	Positives    []int
	Negatives    []int
	Value        struct {
		Standart float64
		Error    float64
	}
}

type Traits []Trait

var traitsInstance *Traits = nil

func TraitsGetInstance(filename string, parser stuff.MarkupParser) *Traits {
	traitsOnce.Do(func() {
		traits := Traits{}
		construct_object(filename, "Traits", parser, &traits)
		traitsInstance = &traits
	})

	return traitsInstance
}
