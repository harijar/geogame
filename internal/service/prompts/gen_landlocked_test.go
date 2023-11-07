package prompts

import (
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var casesGenLandlocked = []struct {
	country *countries.Country
	prev    []*Prompt
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, Landlocked: true}, []*Prompt{},
		&Prompt{ID: LandlockedID, Text: "This country is landlocked"}},
	{&countries.Country{ID: 1, Island: true, Landlocked: false}, []*Prompt{{ID: IslandID}}, nil},
}

func TestPrompts_GenLandlocked(t *testing.T) {
	p := &Prompts{}
	for index, cs := range casesGenLandlocked {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, _ := p.Gen(LandlockedID, cs.country, cs.prev)

			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
