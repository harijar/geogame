package prompts

import (
	repo "geogame/internal/repo/languages"
	"math/rand"
	"strings"
)

func getLanguage(id int, allLanguages repo.Languages, data promptData) string {
	if languages, ok := allLanguages.Languages[id]; ok {
		n := rand.Intn(len(languages))

		// If the country is Sweden and the prompt says "official language is Swedish", it's pretty obvious what country is.
		// So this constraint makes the game more difficult and interesting.
		if strings.Contains(languages[n].Language, data.Country[:3]) {
			return ""
		}
		return languages[n].Language + " is an official language of this country"
	} else {
		return ""
	}
}
