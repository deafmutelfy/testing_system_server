package model

type TraitModelEntry struct {
	Trait   *Trait
	Counter int
}

type TraitModel struct {
	Total   int
	Entries []TraitModelEntry
}

func TraitModelNew() *TraitModel {
	model := TraitModel{}

	t := TraitsGetInstance("", nil)
	for k, _ := range *t {
		v := (*t)[k]
		entry := TraitModelEntry{}
		entry.Trait = &v

		model.Entries = append(model.Entries, entry)
	}

	return &model
}
