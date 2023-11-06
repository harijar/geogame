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
	{MonarchyID, &countries.Country{Monarchy: true}, []*Prompt{},
		&Prompt{ID: MonarchyID, Text: "This country is a monarchy"}, nil},
	{GDPID, &countries.Country{GDP: 228}, []*Prompt{},
		&Prompt{ID: GDPID, Text: "GDP of this country is 228 USD"}, nil},
	{ReligionID,
		&countries.Country{Religion: "Rastafarianism", ReligionPerc: 420}, []*Prompt{},
		&Prompt{ID: ReligionID, Text: "Major religion of this country is Rastafarianism. 420&#37; of people there practice it."}, nil},
	{LanguageID,
		&countries.Country{Languages: []*countries.Language{{Name: "Zalupa-Congolese"}}}, []*Prompt{},
		&Prompt{ID: LanguageID, Text: "Official language of this country is Zalupa-Congolese"}, nil},
	{LanguageID, &countries.Country{}, []*Prompt{}, nil, nil},
	{CapitalID, &countries.Country{}, []*Prompt{}, nil, nil},
	{100, &countries.Country{}, []*Prompt{}, nil, fmt.Errorf("prompt ID not correct")},
}

// Cases for static prompt generation only (so prev is empty here)
var casesGenRandom = []struct {
	country *countries.Country
	prev    []*Prompt
	prompt  *Prompt
	error   error
}{
	{&countries.Country{Monarchy: true},
		[]*Prompt{{ID: 1}, {ID: 3}}, &Prompt{ID: MonarchyID, Text: "This country is a monarchy"}, nil},
	{&countries.Country{Monarchy: true, GDP: 228},
		[]*Prompt{{ID: 1}, {ID: MonarchyID}, {ID: 3}}, &Prompt{ID: GDPID, Text: "GDP of this country is 228 USD"}, nil},
	{&countries.Country{},
		[]*Prompt{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10}, {ID: 11}, {ID: 12}, {ID: 13}, {ID: 14}, {ID: 15}, {ID: 16}},
		nil, fmt.Errorf("unable to find prompt")},
}

func TestPrompts_Gen(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)
	// GetAnotherRandom(&countries.Country{ID: 1}) will return &countries.Country{ID: 2}
	// keep in mind that the argument is a struct, so you need to call this with same pointer as I understand
	countriesRepo.EXPECT().GetAnotherRandom(&countries.Country{ID: 1}).Return(&countries.Country{ID: 2})
	// Get(1) will return &countries.Country{ID: 1}
	// by default the specified data will be returned once, but you can control this behaviour
	countriesRepo.EXPECT().Get(1).Return(&countries.Country{ID: 1}).AnyTimes() // .Times(3) .MaxTimes(10) .MinTimes(2)
	/*
		country := countriesRepo.Get(1)
		t.Log(country.ID)
	*/

	//countriesRepo := &countries.Countries{cache: []*countries.Country{
	//	{ID: 1, Name: "Russian Empire"},
	//	{ID: 2, Name: "Mongol Empire"},
	//}}
	p := &Prompts{countriesRepo: countriesRepo}

	for index, cs := range casesGen {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, err := p.Gen(cs.id, cs.country, cs.prev)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}

func TestPrompts_GenRandom(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//mockCountries := mocks.NewMockCountries(ctrl)

	//countriesRepo := &countries.Countries{cache: []*countries.Country{
	//	{ID: 1, Name: "Russian Empire"},
	//	{ID: 2, Name: "Mongol Empire"},
	//}}
	p := &Prompts{countriesRepo: countriesRepo}
	for index, cs := range casesGenRandom {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			prompt, err := p.GenRandom(cs.country, cs.prev)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
