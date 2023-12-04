package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
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
	{IndependentID, &countries.Country{ID: 1, IndependentFrom: "Antarctica"}, []*Prompt{},
		&Prompt{ID: IndependentID, Text: "This country used to be part of Antarctica"}, nil},
	{ReligionID,
		&countries.Country{ID: 1, Religion: "Rastafarianism", ReligionPerc: 420}, []*Prompt{},
		&Prompt{ID: ReligionID, Text: "Major religion of this country is Rastafarianism. 420% of people there practice it."}, nil},
	{LanguageID,
		&countries.Country{ID: 1, Languages: []*countries.Language{{Name: "Zalupa-Congolese"}}}, []*Prompt{},
		&Prompt{ID: LanguageID, Text: "Official language of this country is Zalupa-Congolese"}, nil},
	{LanguageID, &countries.Country{ID: 1}, []*Prompt{}, nil, nil},
	{CapitalID, &countries.Country{ID: 1}, []*Prompt{}, nil, nil},
	{100, &countries.Country{ID: 1}, []*Prompt{}, nil, fmt.Errorf("prompt ID not correct")},
}

func TestPrompts_Gen(t *testing.T) {
	p := &Prompts{}
	p.logger = zap.Must(zap.NewProduction())
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
		[]*Prompt{{ID: 1}, {ID: 3}, {ID: HemisphereLatID}, {ID: HemisphereLongID}, {ID: LandlockedID}, {ID: IslandID}}, &Prompt{ID: MonarchyID, Text: "This country is a monarchy"}, nil},
	{&countries.Country{ID: 1, Monarchy: true, Funfacts: []*countries.Funfact{{CountryID: 1, Text: "psa popierdolilo"}}},
		[]*Prompt{{ID: 1}, {ID: MonarchyID}, {ID: 3}, {ID: HemisphereLatID}, {ID: HemisphereLongID}, {ID: LandlockedID}, {ID: IslandID}},
		&Prompt{ID: FunfactID, Text: "psa popierdolilo"}, nil},
	{&countries.Country{ID: 1},
		[]*Prompt{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10}, {ID: 11}, {ID: 12}, {ID: 13}, {ID: 14}, {ID: 15}, {ID: 16}, {ID: 17}, {ID: 18}, {ID: 21}, {ID: 22}, {ID: 23}, {ID: 24}, {ID: 25}, {ID: 26}, {ID: 27}},
		nil, fmt.Errorf("failed to find prompt")},
}

func TestPrompts_GenRandom(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	countriesRepo := mocks.NewMockCountries(mockCtrl)

	p := &Prompts{countriesRepo: countriesRepo}
	p.logger = zap.Must(zap.NewProduction())
	for index, cs := range casesGenRandom {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(&countries.Country{ID: 2}).AnyTimes()
			prompt, err := p.GenRandom(cs.country, cs.prev)

			assert.Equal(t, cs.error, err)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
