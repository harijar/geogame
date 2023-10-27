package prompts

import (
	"fmt"
	countries "geogame/internal/repo/countries"
)

func getArea(country *countries.Country) (string, error) {
	if country.Area != 0 {
		return fmt.Sprintf("Area of this country is %v km²", country.Area), nil
	}
	return "", nil
}

func getPopulation(country *countries.Country) (string, error) {
	if country.Population != 0 {
		return fmt.Sprintf("Population of this country is %v people", country.Population), nil
	}
	return "", nil
}

func getGDP(country *countries.Country) (string, error) {
	if country.GDP != 0 {
		return fmt.Sprintf("GDP of this country is %v USD", country.GDP), nil
	}
	return "", nil
}

func getGDPPerCapita(country *countries.Country) (string, error) {
	if country.GDPPerCapita != 0 {
		return fmt.Sprintf("GDP per capita of this country is %v USD", country.GDPPerCapita), nil
	}
	return "", nil
}

func getHDI(country countries.Country) (string, error) {
	if country.HDI != 0 {
		return fmt.Sprintf("HDI of this country is %v", country.HDI), nil
	}
	return "", nil
}

func getArgicultural(country countries.Country) (string, error) {
	if country.AgriculturalSector != 0 {
		return fmt.Sprintf("Argicultural sector of this country is %v &#37; of its GDP", country.AgriculturalSector), nil
	}
	return "", nil
}

func getIndustrial(country countries.Country) (string, error) {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Industrial sector of this country is %v &#37; of its GDP", country.IndustrialSector), nil
	}
	return "", nil
}

func getService(country countries.Country) (string, error) {
	if country.IndustrialSector != 0 {
		return fmt.Sprintf("Service sector of this country is %v &#37; of its GDP", country.ServiceSector), nil
	}
	return "", nil
}
