package prompts

import (
	"github.com/harijar/geogame/internal/mocks"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"strconv"
	"testing"
)

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
	p.logger = zap.Must(zap.NewProduction())
	for index, cs := range casesHemisphereLat {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			if len(cs.prev) > 0 {
				countriesRepo.EXPECT().Get(cs.prev[0].AnotherCountryID).Return(cs.another)
			}
			prompt, _ := p.Gen(HemisphereLatID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
	for index, cs := range casesHemisphereLong {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			if len(cs.prev) > 0 {
				countriesRepo.EXPECT().Get(cs.prev[0].AnotherCountryID).Return(cs.another)
			}
			prompt, _ := p.Gen(HemisphereLongID, cs.country, cs.prev)
			assert.Equal(t, cs.prompt, prompt)
		})
	}
}
