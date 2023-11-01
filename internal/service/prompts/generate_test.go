package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var casesGen = []struct {
	id      int
	country *countries.Country
	prompt  string
	error   error
}{
	{MonarchyID, &countries.Country{Monarchy: true}, "This country is a monarchy", nil},
	{GDPID, &countries.Country{GDP: 228}, "GDP of this country is 228 USD", nil},
	{ReligionID,
		&countries.Country{Religion: "Rastafarianism", ReligionPerc: 420},
		"Major religion of this country is Rastafarianism. 420&#37; of people there practice it.", nil},
	{LanguageID,
		&countries.Country{Languages: []*countries.Language{{Name: "Zalupa-Congolese"}}},
		"Official language of this country is Zalupa-Congolese", nil},
	{LanguageID, &countries.Country{}, "", nil},
	{CapitalID, &countries.Country{}, "", nil},
	{19, &countries.Country{}, "", fmt.Errorf("prompt ID not correct")},
}

var casesGenRandom = []struct {
	country *countries.Country
	prev    []int
	id      int
	prompt  string
	error   error
}{
	{&countries.Country{Monarchy: true},
		[]int{1, 3}, MonarchyID, "This country is a monarchy", nil},
	{&countries.Country{Monarchy: true, GDP: 228},
		[]int{1, 2, 3}, GDPID, "GDP of this country is 228 USD", nil},
	{&countries.Country{},
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}, -1, "", fmt.Errorf("unable to find prompt")},
}

func TestPrompts_Gen(t *testing.T) {
	p := &Prompts{}
	for index, cs := range casesGen {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, err := p.Gen(cs.id, cs.country)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}

func TestPrompts_GenRandom(t *testing.T) {
	p := &Prompts{}
	for index, cs := range casesGenRandom {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			id, prompt, err := p.GenRandom(cs.country, cs.prev)

			assert.Equal(t, cs.id, id)
			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
