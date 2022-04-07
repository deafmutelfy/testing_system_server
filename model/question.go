package model

import (
	"sync"

	"github.com/vokestd/TestingSystem-server/stuff"
)

var questionsOnce sync.Once

type Question struct {
	Id    int
	Title string
}

type Questions []Question

var questionsInstance *Questions = nil

func QuestionsGetInstance(filename string, parser stuff.MarkupParser) *Questions {
	questionsOnce.Do(func() {
		questions := Questions{}
		construct_object(filename, "Questions", parser, &questions)
		questionsInstance = &questions
	})

	return questionsInstance
}

func (q *Questions) GetById(id int) *Question {
	for k, _ := range *q {
		v := (*q)[k]

		if v.Id == id {
			return &v
		}
	}

	return nil
}
