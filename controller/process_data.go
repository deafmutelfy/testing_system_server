package controller

import (
	"fmt"

	"github.com/vokestd/TestingSystem-server/model"
)

const (
	Yes      = iota + 1 // В основном верно
	Probably            // Отчасти верно
	No                  // Нет
	Nothing             // Не могу решить || Затрудняюсь ответить
)

func ProcessData(data *model.SubmitDataModel) {
	t := model.TraitModelNew()

	for _, v := range *data {
		for k := range t.Entries {
			x := &t.Entries[k]
			delta := 0
			if (is_entry(&x.Trait.Positives, v.Id) && v.Value == Yes) || (is_entry(&x.Trait.Negatives, v.Id) && v.Value == No) {
				delta = 2
			} else if is_entry(&x.Trait.Negatives, v.Id) || is_entry(&x.Trait.Positives, v.Id) {
				if v.Value == Probably {
					delta = 1
				} else if v.Value == Nothing {
					delta = -1
				}
			}

			x.Counter += delta
			t.Total += delta
		}
	}

	fmt.Println(t)
}
