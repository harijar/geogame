package prompts

import (
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"strconv"
	"testing"
)

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
	for index, cs := range casesGenLocationLat {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.another)
			prompt, _ := p.Gen(LocationLatID, cs.country, cs.prev)

			assert.Equal(t, cs.prompt, prompt)
		})
	}

	for index, cs := range casesGenLocationLong {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			countriesRepo.EXPECT().GetAnotherRandom(cs.country).Return(cs.another)
			prompt, _ := p.Gen(LocationLongID, cs.country, cs.prev)

			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
