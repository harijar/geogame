package prompts

import (
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
)

type caseNumeric struct {
	country *countries.Country
	place   int
	prompts [2]*Prompt
}

var casesGenArea = []caseNumeric{
	{&countries.Country{ID: 1, Area: 420}, 1,
		[2]*Prompt{{ID: AreaID, Text: "Area of this country is 420 km²"},
			{ID: AreaID, Text: "This country is number 1 in terms of area"}}},
	{&countries.Country{ID: 1}, 0, [2]*Prompt{nil, nil}},
}

var casesGenPopulation = []caseNumeric{
	{&countries.Country{ID: 1, Population: 69}, 2,
		[2]*Prompt{{ID: PopulationID, Text: "Population of this country is 69 people"},
			{ID: PopulationID, Text: "This country is number 2 in terms of population"}}},
	{&countries.Country{ID: 1}, 0, [2]*Prompt{nil, nil}},
}

var casesGenGDP = []caseNumeric{
	{&countries.Country{ID: 1, GDP: 1488}, 3,
		[2]*Prompt{{ID: GDPID, Text: "GDP of this country is 1 488 million USD"},
			{ID: GDPID, Text: "This country is number 3 in terms of GDP"}}},
	{&countries.Country{ID: 1}, 0, [2]*Prompt{nil, nil}},
}

var casesGenGDPPerCapita = []caseNumeric{
	{&countries.Country{ID: 1, GDPPerCapita: 777}, 4,
		[2]*Prompt{{ID: GDPPerCapitaID, Text: "GDP per capita of this country is 777 USD"},
			{ID: GDPPerCapitaID, Text: "This country is number 4 in terms of GDP per capita"}}},
	{&countries.Country{ID: 1}, 0, [2]*Prompt{nil, nil}},
}

var casesGenHDI = []caseNumeric{
	{&countries.Country{ID: 1, HDI: 0.666}, 5,
		[2]*Prompt{{ID: HDIID, Text: "HDI of this country is 0.666"},
			{ID: HDIID, Text: "This country is number 5 in terms of HDI"}}},
	{&countries.Country{ID: 1}, 0, [2]*Prompt{nil, nil}},
}

func TestPrompts_GenNumeric(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := Prompts{}
	p.countriesRepo = countriesRepo
	for index, cs := range casesGenArea {
		t.Run("area_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceArea(cs.country).Return(cs.place).AnyTimes()
			prompt, _ := p.Gen(AreaID, cs.country, nil)
			if prompt == nil {
				assert.Equal(t, cs.prompts[0], prompt)
			} else {
				if *prompt != *cs.prompts[0] {
					assert.Equal(t, cs.prompts[1], prompt)
				} else {
					assert.Equal(t, cs.prompts[0], prompt)
				}
			}
		})
	}
	for index, cs := range casesGenPopulation {
		t.Run("population_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlacePopulation(cs.country).Return(cs.place).AnyTimes()
			prompt, _ := p.Gen(PopulationID, cs.country, nil)
			if prompt == nil {
				assert.Equal(t, cs.prompts[0], prompt)
			} else {
				if *prompt != *cs.prompts[0] {
					assert.Equal(t, cs.prompts[1], prompt)
				} else {
					assert.Equal(t, cs.prompts[0], prompt)
				}
			}
		})
	}
	for index, cs := range casesGenGDP {
		t.Run("gdp_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceGDP(cs.country).Return(cs.place).AnyTimes()
			prompt, _ := p.Gen(GDPID, cs.country, nil)
			if prompt == nil {
				assert.Equal(t, cs.prompts[0], prompt)
			} else {
				if *prompt != *cs.prompts[0] {
					assert.Equal(t, cs.prompts[1], prompt)
				} else {
					assert.Equal(t, cs.prompts[0], prompt)
				}
			}
		})
	}
	for index, cs := range casesGenGDPPerCapita {
		t.Run("gdp_per_capita_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceGDPPerCapita(cs.country).Return(cs.place).AnyTimes()
			prompt, _ := p.Gen(GDPPerCapitaID, cs.country, nil)
			if prompt == nil {
				assert.Equal(t, cs.prompts[0], prompt)
			} else {
				if *prompt != *cs.prompts[0] {
					assert.Equal(t, cs.prompts[1], prompt)
				} else {
					assert.Equal(t, cs.prompts[0], prompt)
				}
			}
		})
	}
	for index, cs := range casesGenHDI {
		t.Run("hdi_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceHDI(cs.country).Return(cs.place).AnyTimes()
			prompt, _ := p.Gen(HDIID, cs.country, nil)
			if prompt == nil {
				assert.Equal(t, cs.prompts[0], prompt)
			} else {
				if *prompt != *cs.prompts[0] {
					assert.Equal(t, cs.prompts[1], prompt)
				} else {
					assert.Equal(t, cs.prompts[0], prompt)
				}
			}
		})
	}
}
