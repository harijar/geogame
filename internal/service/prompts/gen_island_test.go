package prompts

import (
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var casesGenIsland = []struct {
	country *countries.Country
	prev    []*Prompt
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, Island: true}, []*Prompt{},
		&Prompt{ID: IslandID, Text: "This country is an island country"}},
	{&countries.Country{ID: 1, Island: false, Landlocked: true}, []*Prompt{{ID: LandlockedID}}, nil},
}

func TestPrompts_GenIsland(t *testing.T) {
	p := &Prompts{}
	for index, cs := range casesGenIsland {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, _ := p.Gen(IslandID, cs.country, cs.prev)

			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
