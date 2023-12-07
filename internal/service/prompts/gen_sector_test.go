package prompts

import (
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var casesGenSector = []struct {
	promptID int
	country  *countries.Country
	prev     []*Prompt
	prompt   *Prompt
}{
	{AgriculturalSectorID, &countries.Country{ID: 1, AgriculturalSector: 66},
		nil, &Prompt{ID: AgriculturalSectorID, Text: "Agricultural sector of this country is 66% of its GDP"}},
	{AgriculturalSectorID, &countries.Country{ID: 1, AgriculturalSector: 66},
		[]*Prompt{{ID: IndustrialSectorID}, {ID: ServiceSectorID}}, nil},
	{IndustrialSectorID, &countries.Country{ID: 1, IndustrialSector: 66},
		nil, &Prompt{ID: IndustrialSectorID, Text: "Industrial sector of this country is 66% of its GDP"}},
	{IndustrialSectorID, &countries.Country{ID: 1, AgriculturalSector: 66},
		nil, nil},
	{ServiceSectorID, &countries.Country{ID: 1, ServiceSector: 66},
		nil, &Prompt{ID: ServiceSectorID, Text: "Service sector of this country is 66% of its GDP"}},
	{ServiceSectorID, &countries.Country{ID: 1, ServiceSector: 66},
		[]*Prompt{{ID: ServiceSectorID}}, nil},
}

func TestPrompts_GenSector(t *testing.T) {
	p := &Prompts{}
	for index, cs := range casesGenSector {
		t.Run("sector_"+strconv.Itoa(index), func(t *testing.T) {
			prompt, _ := p.Gen(cs.promptID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
