package prompts

import (
	repo "geogame/internal/repo/funfacts"
	"math/rand"
)

func GetFunFact(id int, allFunfacts repo.Funfacts, data *promptData) string {
	if funfacts, ok := allFunfacts.Funtacts[id]; ok {
		n := rand.Intn(len(funfacts))
		return funfacts[n].Funfact
	} else {
		return ""
	}
}
