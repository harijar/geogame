package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
)

// Cases for static prompt generation only (so prev is empty here)
var casesGen = []struct {
	id      int
	country *countries.Country
	prev    []*Prompt
	prompt  *Prompt
	error   error
}{
	{MonarchyID, &countries.Country{ID: 1, Monarchy: true}, []*Prompt{},
		&Prompt{ID: MonarchyID, Text: "This country is a monarchy"}, nil},
	{GDPID, &countries.Country{ID: 1, GDP: 228}, []*Prompt{},
		&Prompt{ID: GDPID, Text: "GDP of this country is 228 USD"}, nil},
	{ReligionID,
		&countries.Country{ID: 1, Religion: "Rastafarianism", ReligionPerc: 420}, []*Prompt{},
		&Prompt{ID: ReligionID, Text: "Major religion of this country is Rastafarianism. 420&#37; of people there practice it."}, nil},
	{LanguageID,
		&countries.Country{ID: 1, Languages: []*countries.Language{{Name: "Zalupa-Congolese"}}}, []*Prompt{},
		&Prompt{ID: LanguageID, Text: "Official language of this country is Zalupa-Congolese"}, nil},
	{LanguageID, &countries.Country{ID: 1}, []*Prompt{}, nil, nil},
	{CapitalID, &countries.Country{ID: 1}, []*Prompt{}, nil, nil},
	{100, &countries.Country{ID: 1}, []*Prompt{}, nil, fmt.Errorf("prompt ID not correct")},
}

func TestPrompts_Gen(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := &Prompts{countriesRepo: countriesRepo}
	for index, cs := range casesGen {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, err := p.Gen(cs.id, cs.country, cs.prev)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}

// Cases for static prompt generation only
var casesGenRandom = []struct {
	country *countries.Country
	prev    []*Prompt
	prompt  *Prompt
	error   error
}{
	{&countries.Country{ID: 1, Monarchy: true},
		[]*Prompt{{ID: 1}, {ID: 3}, {ID: 17}, {ID: 18}}, &Prompt{ID: MonarchyID, Text: "This country is a monarchy"}, nil},
	{&countries.Country{ID: 1, Monarchy: true, GDP: 228},
		[]*Prompt{{ID: 1}, {ID: MonarchyID}, {ID: 3}, {ID: 17}, {ID: 18}}, &Prompt{ID: GDPID, Text: "GDP of this country is 228 USD"}, nil},
	{&countries.Country{ID: 1},
		[]*Prompt{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10}, {ID: 11}, {ID: 12}, {ID: 13}, {ID: 14}, {ID: 15}, {ID: 16}, {ID: 17}, {ID: 18}},
		nil, fmt.Errorf("failed to find prompt")},
}

func TestPrompts_GenRandom(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := &Prompts{countriesRepo: countriesRepo}
	for index, cs := range casesGenRandom {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(&countries.Country{ID: 2}).AnyTimes()
			prompt, err := p.GenRandom(cs.country, cs.prev)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}

var casesGenLocationLat = []struct {
	country *countries.Country
	another *countries.Country
	prev    []*Prompt
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, Northernmost: 45, Southernmost: 35},
		&countries.Country{ID: 2, Name: "Qin Empire", Northernmost: 30, Southernmost: 0},
		[]*Prompt{{ID: HemisphereLatID}}, &Prompt{ID: LocationLatID, Text: "This country is located north to Qin Empire", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Northernmost: -2, Southernmost: -12, HemisphereLat: countries.Southern},
		&countries.Country{ID: 2, Name: "Mongol Empire", Northernmost: 45, Southernmost: 35, HemisphereLat: countries.Northern},
		[]*Prompt{}, &Prompt{ID: LocationLatID, Text: "This country is located south to Mongol Empire", AnotherCountryID: 2}},

	{&countries.Country{ID: 1, Northernmost: -2, Southernmost: -12, HemisphereLat: countries.Southern},
		&countries.Country{ID: 2, Name: "Mongol Empire", Northernmost: 45, Southernmost: 35, HemisphereLat: countries.Northern},
		[]*Prompt{{ID: HemisphereLatID}}, nil},
	{&countries.Country{ID: 1, Northernmost: -10, Southernmost: -16, HemisphereLat: countries.Southern},
		&countries.Country{ID: 2, Northernmost: 16, Southernmost: 10, HemisphereLat: countries.Northern},
		[]*Prompt{{ID: HemisphereLatID}}, nil},

	{&countries.Country{ID: 1, Northernmost: 12, Southernmost: -100},
		&countries.Country{ID: 2, Northernmost: -10, Southernmost: -20},
		[]*Prompt{}, nil},
	{&countries.Country{ID: 1, Northernmost: 12, Southernmost: -12},
		&countries.Country{ID: 2, Northernmost: 11, Southernmost: -13},
		[]*Prompt{}, nil},
}

var casesGenLocationLong = []struct {
	country *countries.Country
	another *countries.Country
	prev    []*Prompt
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, Easternmost: 45, Westernmost: 35},
		&countries.Country{ID: 2, Name: "Roman Empire", Easternmost: 30, Westernmost: 0},
		[]*Prompt{{ID: HemisphereLongID}}, &Prompt{ID: LocationLongID, Text: "This country is located east to Roman Empire", AnotherCountryID: 2}},
	{&countries.Country{ID: 1, Easternmost: -2, Westernmost: -12, HemisphereLong: countries.Western},
		&countries.Country{ID: 2, Name: "Kyivan Rus'", Easternmost: 45, Westernmost: 35, HemisphereLong: countries.Eastern},
		[]*Prompt{}, &Prompt{ID: LocationLongID, Text: "This country is located west to Kyivan Rus'", AnotherCountryID: 2}},

	{&countries.Country{ID: 1, Easternmost: -2, Westernmost: -12, HemisphereLong: countries.Western},
		&countries.Country{ID: 2, Name: "Kyivan Rus'", Easternmost: 45, Westernmost: 35, HemisphereLong: countries.Eastern},
		[]*Prompt{{ID: HemisphereLongID}}, nil},
	{&countries.Country{ID: 1, Easternmost: -10, Westernmost: -16, HemisphereLong: countries.Western},
		&countries.Country{ID: 2, Easternmost: 16, Westernmost: 10, HemisphereLong: countries.Eastern},
		[]*Prompt{{ID: HemisphereLongID}}, nil},

	{&countries.Country{ID: 1, Easternmost: 12, Westernmost: -100},
		&countries.Country{ID: 2, Easternmost: -10, Westernmost: -20},
		[]*Prompt{}, nil},
	{&countries.Country{ID: 1, Easternmost: 12, Westernmost: -12},
		&countries.Country{ID: 2, Easternmost: 11, Westernmost: -13},
		[]*Prompt{}, nil},
}

func TestPrompts_GenLocation(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := &Prompts{countriesRepo: countriesRepo}
	for _, cs := range casesGenLocationLat {
		countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.another)
		prompt, _ := p.Gen(LocationLatID, cs.country, cs.prev)

		assert.Equal(t, cs.prompt, prompt)
	}

	for _, cs := range casesGenLocationLong {
		countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.another)
		prompt, _ := p.Gen(LocationLongID, cs.country, cs.prev)

		assert.Equal(t, cs.prompt, prompt)
	}
}

var casesHemisphereLat = []struct {
	country *countries.Country
	prev    []*Prompt
	another *countries.Country
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, HemisphereLat: countries.Northern},
		[]*Prompt{}, nil, &Prompt{ID: HemisphereLatID, Text: "This country is located in Northern hemisphere"}},
	{&countries.Country{ID: 1, HemisphereLat: countries.Northern, Southernmost: 5},
		[]*Prompt{{ID: LocationLatID, AnotherCountryID: 2}}, &countries.Country{ID: 2, HemisphereLat: countries.Northern, Northernmost: 10},
		&Prompt{ID: HemisphereLatID, Text: "This country is located in Northern hemisphere"}},
	{&countries.Country{ID: 1, HemisphereLat: countries.Northern, Southernmost: 10},
		[]*Prompt{{ID: LocationLatID, AnotherCountryID: 2}}, &countries.Country{ID: 2, HemisphereLat: countries.Northern, Northernmost: 5}, nil},
}

var casesHemisphereLong = []struct {
	country *countries.Country
	prev    []*Prompt
	another *countries.Country
	prompt  *Prompt
}{
	{&countries.Country{ID: 1, HemisphereLong: countries.Eastern},
		[]*Prompt{}, nil, &Prompt{ID: HemisphereLongID, Text: "This country is located in Eastern hemisphere"}},
	{&countries.Country{ID: 1, HemisphereLat: countries.Eastern, Westernmost: 5},
		[]*Prompt{{ID: LocationLongID, AnotherCountryID: 2}}, &countries.Country{ID: 2, HemisphereLong: countries.Eastern, Easternmost: 10},
		&Prompt{ID: HemisphereLongID, Text: "This country is located in Eastern hemisphere"}},
	{&countries.Country{ID: 1, HemisphereLong: countries.Eastern, Westernmost: 10},
		[]*Prompt{{ID: LocationLongID, AnotherCountryID: 2}}, &countries.Country{ID: 2, HemisphereLat: countries.Eastern, Easternmost: 5}, nil},
}

func TestPrompts_Hemisphere(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := &Prompts{countriesRepo: countriesRepo}
	for _, cs := range casesHemisphereLat {
		if len(cs.prev) > 0 {
			countriesRepo.EXPECT().Get(cs.prev[0].AnotherCountryID).Return(cs.another)
		}
		prompt, _ := p.Gen(HemisphereLatID, cs.country, cs.prev)

		assert.Equal(t, cs.prompt, prompt)
	}
	for _, cs := range casesHemisphereLong {
		if len(cs.prev) > 0 {
			countriesRepo.EXPECT().Get(cs.prev[0].AnotherCountryID).Return(cs.another)
		}
		prompt, _ := p.Gen(HemisphereLongID, cs.country, cs.prev)

		assert.Equal(t, cs.prompt, prompt)
	}
}
