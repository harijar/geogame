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
	prompt        *Prompt
}

var casesGenCompareArea = []caseNumericCompare{
	{&countries.Country{ID: 1, Area: 10}, 180,
		[]*countries.Country{{ID: 2, Area: 10000, Name: "Hyperborea"}}, []int{12},
		nil, &Prompt{ID: CompareAreaID, Text: "Area of this country is smaller than that of Hyperborea", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Area: 10}, 180,
		[]*countries.Country{{ID: 2, Area: 10000, Name: "Hyperborea"}}, []int{12},
		[]*Prompt{{ID: AreaID}}, nil},

	{&countries.Country{ID: 1, Area: 10000}, 12,
		[]*countries.Country{{ID: 2, Area: 60000, Name: "Tartaria"}, {ID: 3, Area: 59000, Name: "Atlantida"}}, []int{1, 11},
		nil, &Prompt{ID: CompareAreaID, Text: "Area of this country is smaller than that of Atlantida", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, Area: 10000}, 12,
		[]*countries.Country{{ID: 2, Area: 20, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: AreaID}, {ID: CompareAreaID, AnotherCountryID: 3}}, &Prompt{ID: CompareAreaID, Text: "Area of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Area: 60000}, 1,
		[]*countries.Country{{ID: 2, Area: 20, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Chechnya"}}, []int{170, 168},
		[]*Prompt{{ID: CompareAreaID, AnotherCountryID: 2}}, &Prompt{ID: CompareAreaID, Text: "Area of this country is bigger than that of Chechnya", AnotherCountryID: 3}},
}

var casesGenComparePopulation = []caseNumericCompare{
	{&countries.Country{ID: 1, Population: 10}, 180,
		[]*countries.Country{{ID: 2, Population: 10000, Name: "Hyperborea"}}, []int{12},
		nil, &Prompt{ID: ComparePopulationID, Text: "Population of this country is smaller than that of Hyperborea", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Population: 10}, 180,
		[]*countries.Country{{ID: 2, Population: 10000, Name: "Hyperborea"}}, []int{12},
		[]*Prompt{{ID: PopulationID}}, nil},

	{&countries.Country{ID: 1, Population: 10000}, 12,
		[]*countries.Country{{ID: 2, Population: 60000, Name: "Tartaria"}, {ID: 3, Population: 59000, Name: "Atlantida"}}, []int{1, 11},
		nil, &Prompt{ID: ComparePopulationID, Text: "Population of this country is smaller than that of Atlantida", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, Population: 10000}, 12,
		[]*countries.Country{{ID: 2, Population: 20, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: PopulationID}, {ID: ComparePopulationID, AnotherCountryID: 3}}, &Prompt{ID: ComparePopulationID, Text: "Population of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Population: 60000}, 1,
		[]*countries.Country{{ID: 2, Population: 20, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Chechnya"}}, []int{170, 168},
		[]*Prompt{{ID: ComparePopulationID, AnotherCountryID: 2}}, &Prompt{ID: ComparePopulationID, Text: "Population of this country is bigger than that of Chechnya", AnotherCountryID: 3}},
}

var casesGenCompareGDP = []caseNumericCompare{
	{&countries.Country{ID: 1, GDP: 10}, 180,
		[]*countries.Country{{ID: 2, GDP: 10000, Name: "Hyperborea"}}, []int{12},
		nil, &Prompt{ID: CompareGDPID, Text: "GDP of this country is smaller than that of Hyperborea", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, GDP: 10}, 180,
		[]*countries.Country{{ID: 2, GDP: 10000, Name: "Hyperborea"}}, []int{12},
		[]*Prompt{{ID: GDPID}}, nil},

	{&countries.Country{ID: 1, GDP: 10000}, 12,
		[]*countries.Country{{ID: 2, GDP: 60000, Name: "Tartaria"}, {ID: 3, GDP: 59000, Name: "Atlantida"}}, []int{1, 11},
		nil, &Prompt{ID: CompareGDPID, Text: "GDP of this country is smaller than that of Atlantida", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, GDP: 10000}, 12,
		[]*countries.Country{{ID: 2, GDP: 20, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: GDPID}, {ID: CompareGDPID, AnotherCountryID: 3}}, &Prompt{ID: CompareGDPID, Text: "GDP of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, GDP: 60000}, 1,
		[]*countries.Country{{ID: 2, GDP: 20, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Chechnya"}}, []int{170, 168},
		[]*Prompt{{ID: CompareGDPID, AnotherCountryID: 2}}, &Prompt{ID: CompareGDPID, Text: "GDP of this country is bigger than that of Chechnya", AnotherCountryID: 3}},
}

var casesGenCompareGDPPerCapita = []caseNumericCompare{
	{&countries.Country{ID: 1, GDPPerCapita: 10}, 180,
		[]*countries.Country{{ID: 2, GDPPerCapita: 10000, Name: "Hyperborea"}}, []int{12},
		nil, &Prompt{ID: CompareGDPPerCapitaID, Text: "GDP per capita of this country is smaller than that of Hyperborea", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, GDPPerCapita: 10}, 180,
		[]*countries.Country{{ID: 2, GDPPerCapita: 10000, Name: "Hyperborea"}}, []int{12},
		[]*Prompt{{ID: GDPPerCapitaID}}, nil},

	{&countries.Country{ID: 1, GDPPerCapita: 10000}, 12,
		[]*countries.Country{{ID: 2, GDPPerCapita: 60000, Name: "Tartaria"}, {ID: 3, GDPPerCapita: 59000, Name: "Atlantida"}}, []int{1, 11},
		nil, &Prompt{ID: CompareGDPPerCapitaID, Text: "GDP per capita of this country is smaller than that of Atlantida", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, GDPPerCapita: 10000}, 12,
		[]*countries.Country{{ID: 2, GDPPerCapita: 20, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: GDPPerCapitaID}, {ID: CompareGDPPerCapitaID, AnotherCountryID: 3}}, &Prompt{ID: CompareGDPPerCapitaID, Text: "GDP per capita of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, GDPPerCapita: 60000}, 1,
		[]*countries.Country{{ID: 2, GDPPerCapita: 20, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Chechnya"}}, []int{170, 168},
		[]*Prompt{{ID: CompareGDPPerCapitaID, AnotherCountryID: 2}}, &Prompt{ID: CompareGDPPerCapitaID, Text: "GDP per capita of this country is bigger than that of Chechnya", AnotherCountryID: 3}},
}

var casesGenCompareHDI = []caseNumericCompare{
	{&countries.Country{ID: 1, HDI: 0.3}, 180,
		[]*countries.Country{{ID: 2, HDI: 0.8, Name: "Hyperborea"}}, []int{12},
		nil, &Prompt{ID: CompareHDIID, Text: "HDI of this country is smaller than that of Hyperborea", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, HDI: 0.3}, 180,
		[]*countries.Country{{ID: 2, HDI: 0.8, Name: "Hyperborea"}}, []int{12},
		[]*Prompt{{ID: HDIID}}, nil},

	{&countries.Country{ID: 1, HDI: 0.8}, 12,
		[]*countries.Country{{ID: 2, HDI: 0.999, Name: "Tartaria"}, {ID: 3, HDI: 0.95, Name: "Atlantida"}}, []int{1, 11},
		nil, &Prompt{ID: CompareHDIID, Text: "HDI of this country is smaller than that of Atlantida", AnotherCountryID: 3}},
	{&countries.Country{ID: 1, HDI: 0.8}, 12,
		[]*countries.Country{{ID: 2, HDI: 0.6, Name: "Dagestan"}}, []int{170},
		[]*Prompt{{ID: HDIID}, {ID: CompareHDIID, AnotherCountryID: 3}}, &Prompt{ID: CompareHDIID, Text: "HDI of this country is bigger than that of Dagestan", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, HDI: 0.999}, 1,
		[]*countries.Country{{ID: 2, HDI: 0.6, Name: "Dagestan"}, {ID: 3, Area: 21, Name: "Chechnya"}}, []int{170, 168},
		[]*Prompt{{ID: CompareHDIID, AnotherCountryID: 2}}, &Prompt{ID: CompareHDIID, Text: "HDI of this country is bigger than that of Chechnya", AnotherCountryID: 3}},
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
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceArea(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceArea(cs.anothers[1]).Return(cs.anotherPlaces[1])
			}
			prompt, _ := p.Gen(CompareAreaID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenComparePopulation {
		t.Run("compare_population_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlacePopulation(cs.country).Return(cs.place).AnyTimes()
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlacePopulation(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlacePopulation(cs.anothers[1]).Return(cs.anotherPlaces[1])
			}
			prompt, _ := p.Gen(ComparePopulationID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenCompareGDP {
		t.Run("compare_gdp_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceGDP(cs.country).Return(cs.place).AnyTimes()
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceGDP(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceGDP(cs.anothers[1]).Return(cs.anotherPlaces[1])
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
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceGDPPerCapita(cs.anothers[1]).Return(cs.anotherPlaces[1])
			}
			prompt, _ := p.Gen(CompareGDPPerCapitaID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesGenCompareHDI {
		t.Run("compare_hdi_"+strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetPlaceHDI(cs.country).Return(cs.place).AnyTimes()
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[0]).MaxTimes(1)
			countriesRepo.EXPECT().GetPlaceHDI(cs.anothers[0]).Return(cs.anotherPlaces[0]).MaxTimes(1)
			if len(cs.anotherPlaces) > 1 {
				countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.anothers[1])
				countriesRepo.EXPECT().GetPlaceHDI(cs.anothers[1]).Return(cs.anotherPlaces[1])
			}
			prompt, _ := p.Gen(CompareHDIID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
