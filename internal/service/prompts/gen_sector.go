package prompts

import (
	"fmt"
	"github.com/harijar/geogame/internal/repo/postgres/countries"
)

func (p *Prompts) genAgriculturalSector(c *countries.Country, prev []*Prompt) *Prompt {
	count := 0
	for _, prompt := range prev {
		if prompt.ID == AgriculturalSectorID {
			return nil
		} else if prompt.ID == IndustrialSectorID || prompt.ID == ServiceSectorID {
			count++
		}
	}
	if count < 2 && c.AgriculturalSector != 0 {
		return &Prompt{ID: AgriculturalSectorID, Text: fmt.Sprintf("Agricultural sector of this country is %v%% of its GDP", c.AgriculturalSector)}
	}
	return nil
}

func (p *Prompts) genIndustrialSector(c *countries.Country, prev []*Prompt) *Prompt {
	count := 0
	for _, prompt := range prev {
		if prompt.ID == IndustrialSectorID {
			return nil
		} else if prompt.ID == AgriculturalSectorID || prompt.ID == ServiceSectorID {
			count++
		}
	}
	if count < 2 && c.IndustrialSector != 0 {
		return &Prompt{ID: IndustrialSectorID, Text: fmt.Sprintf("Industrial sector of this country is %v%% of its GDP", c.IndustrialSector)}
	}
	return nil
}

func (p *Prompts) genServiceSector(c *countries.Country, prev []*Prompt) *Prompt {
	count := 0
	for _, prompt := range prev {
		if prompt.ID == ServiceSectorID {
			return nil
		} else if prompt.ID == AgriculturalSectorID || prompt.ID == IndustrialSectorID {
			count++
		}
	}
	if count < 2 && c.ServiceSector != 0 {
		return &Prompt{ID: ServiceSectorID, Text: fmt.Sprintf("Service sector of this country is %v%% of its GDP", c.ServiceSector)}
	}
	return nil
}
