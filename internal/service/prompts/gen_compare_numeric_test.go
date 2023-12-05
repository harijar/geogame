package prompts

import (
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
)

type caseNumericCompare struct {
	country       *countries.Country
	place         int
	anothers      []*countries.Country
	anotherPlaces []int
	prev          []*Prompt
	prevCountry   *countries.Country
	prevPlace     int
	prompt        *Prompt
}

var casesGenCompareArea = []caseNumericCompare{
	{&countries.Country{ID: 1, Area: 10}, 180,
		[]*countries.Country{{ID: 2, Area: 10000, Name: "Great Khalifah"}}, []int{12},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: CompareAreaID, Text: "Area of this country is smaller than that of Great Khalifah", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Area: 30}, 100,
		[]*countries.Country{{ID: 2, Area: 10000, Name: "Great Khalifah"}, {ID: 3, Area: 200, Name: "California"}}, []int{12, 70},
		[]*Prompt{{ID: CompareAreaID}}, &countries.Country{ID: 0}, 50,
		&Prompt{ID: CompareAreaID, Text: "Area of this country is smaller than that of California", AnotherCountryID: 3}},

	{&countries.Country{ID: 1, Area: 10000}, 12,
		[]*countries.Country{{ID: 2, Area: 60000, Name: "Tartaria"}, {ID: 3, Area: 59000, Name: "Hyperborea"}}, []int{1, 11},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: CompareAreaID, Text: "Area of this country is smaller than that of Hyperborea", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, Area: 10000}, 12,
		[]*countries.Country{{ID: 2, Area: 20, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: CompareAreaID, AnotherCountryID: 3}}, &countries.Country{ID: 3}, 11,
		&Prompt{ID: CompareAreaID, Text: "Area of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Area: 60000}, 1,
		[]*countries.Country{{ID: 2, Area: 20, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Raqqa"}}, []int{170, 168},
		[]*Prompt{{ID: CompareAreaID, AnotherCountryID: 2}}, &countries.Country{ID: 2}, 171,
		&Prompt{ID: CompareAreaID, Text: "Area of this country is bigger than that of Raqqa", AnotherCountryID: 3}},
}

var casesGenComparePopulation = []caseNumericCompare{
	{&countries.Country{ID: 1, Population: 10}, 180,
		[]*countries.Country{{ID: 2, Population: 10000, Name: "Great Khalifah"}}, []int{12},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: ComparePopulationID, Text: "Population of this country is smaller than that of Great Khalifah", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Population: 30}, 100,
		[]*countries.Country{{ID: 2, Population: 10000, Name: "Great Khalifah"}, {ID: 3, Population: 10, Name: "Amsterdam"}}, []int{12, 130},
		[]*Prompt{{ID: ComparePopulationID}}, &countries.Country{ID: 0}, 50,
		&Prompt{ID: ComparePopulationID, Text: "Population of this country is bigger than that of Amsterdam", AnotherCountryID: 3}},
}

var casesGenCompareGDP = []caseNumericCompare{
	{&countries.Country{ID: 1, GDP: 10}, 180,
		[]*countries.Country{{ID: 2, GDP: 10000, Name: "Great Khalifah"}}, []int{12},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: CompareGDPID, Text: "GDP of this country is smaller than that of Great Khalifah", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, GDP: 30}, 100,
		[]*countries.Country{{ID: 2, GDP: 20, Name: "Barbieland"}, {ID: 3, GDP: 10000, Name: "Great Khalifah"}}, []int{155, 12},
		[]*Prompt{{ID: CompareGDPID}}, &countries.Country{ID: 0}, 150,
		&Prompt{ID: CompareGDPID, Text: "GDP of this country is smaller than that of Great Khalifah", AnotherCountryID: 3}},
}

var casesGenCompareGDPPerCapita = []caseNumericCompare{
	{&countries.Country{ID: 1, GDPPerCapita: 10}, 180,
		[]*countries.Country{{ID: 2, GDPPerCapita: 10000, Name: "Great Khalifah"}}, []int{12},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: CompareGDPPerCapitaID, Text: "GDP per capita of this country is smaller than that of Great Khalifah", AnotherCountryID: 2}},
}

var casesGenCompareHDI = []caseNumericCompare{
	{&countries.Country{ID: 1, HDI: 0.3}, 180,
		[]*countries.Country{{ID: 2, HDI: 0.8, Name: "Great Khalifah"}}, []int{12},
		nil, &countries.Country{ID: 3}, 0,
		&Prompt{ID: CompareHDIID, Text: "HDI of this country is smaller than that of Great Khalifah", AnotherCountryID: 2}},
}

func TestPrompts_GenCompareNumeric(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := Prompts{}
	p.countriesRepo = countriesRepo
	for index, cs := range casesGenCompareArea {
		t.Run("compare_area_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceArea(cs.country).Return(cs.place).AnyTimes()

			countriesRepo.EXPECT().Get(cs.prevCountry.ID).Return(cs.prevCountry).AnyTimes()
			countriesRepo.EXPECT().GetPlaceArea(cs.prevCountry).Return(cs.prevPlace).AnyTimes()

			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceArea(cs.anothers[0]).Return(cs.anotherPlaces[0]).AnyTimes()
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceArea(cs.anothers[1]).Return(cs.anotherPlaces[1]).AnyTimes()
			}
			prompt, _ := p.Gen(CompareAreaID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenComparePopulation {
		t.Run("compare_population_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlacePopulation(cs.country).Return(cs.place).AnyTimes()

			countriesRepo.EXPECT().Get(cs.prevCountry.ID).Return(cs.prevCountry).AnyTimes()
			countriesRepo.EXPECT().GetPlacePopulation(cs.prevCountry).Return(cs.prevPlace).AnyTimes()

			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlacePopulation(cs.anothers[0]).Return(cs.anotherPlaces[0]).AnyTimes()
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlacePopulation(cs.anothers[1]).Return(cs.anotherPlaces[1]).AnyTimes()
			}
			prompt, _ := p.Gen(ComparePopulationID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenCompareGDP {
		t.Run("compare_gdp_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceGDP(cs.country).Return(cs.place).AnyTimes()

			countriesRepo.EXPECT().Get(cs.prevCountry.ID).Return(cs.prevCountry).AnyTimes()
			countriesRepo.EXPECT().GetPlaceGDP(cs.prevCountry).Return(cs.prevPlace).AnyTimes()

			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceGDP(cs.anothers[0]).Return(cs.anotherPlaces[0]).AnyTimes()
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceGDP(cs.anothers[1]).Return(cs.anotherPlaces[1]).AnyTimes()
			}
			prompt, _ := p.Gen(CompareGDPID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenCompareGDPPerCapita {
		t.Run("compare_gdp_per_capita_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceGDPPerCapita(cs.country).Return(cs.place).AnyTimes()
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceGDPPerCapita(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			prompt, _ := p.Gen(CompareGDPPerCapitaID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenCompareHDI {
		t.Run("compare_hdi_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceHDI(cs.country).Return(cs.place).AnyTimes()
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceHDI(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			prompt, _ := p.Gen(CompareHDIID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
