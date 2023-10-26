package prompts

import (
	repo "geogame/internal/repo/countries"
	"strconv"
)

func getArea(country *repo.Country) string {
	if country.Area != 0 {
		area := strconv.FormatFloat(country.Area, 'f', -1, 64)
		return "Area of this country is " + area + " km²"
	} else {
		return ""
	}
}

func getPopulation(country *repo.Country) string {
	if country.Population != 0 {
		return "Population of this country is " + strconv.Itoa(country.Population) + " USD"
	} else {
		return ""
	}
}

func getGDP(country *repo.Country) string {
	if country.GDP != 0 {
		return "GDP of this country is " + strconv.Itoa(country.GDP) + " USD"
	} else {
		return ""
	}
}
func getGDPPerCapita(country *repo.Country) string {
	if country.GDPPerCapita != 0 {
		return "GDP per caita of this country is " + strconv.Itoa(country.GDPPerCapita) + " USD"
	} else {
		return ""
	}
}

func getHDI(country *repo.Country) string {
	if country.HDI != 0 {
		hdi := strconv.FormatFloat(country.HDI, 'f', 3, 64)
		return "HDI of this country is " + hdi
	} else {
		return ""
	}
}

func getArgicultural(country *repo.Country) string {
	if country.AgriculturalSector != 0 {
		agr := strconv.FormatFloat(country.AgriculturalSector, 'f', -1, 64)
		return "Agricultural sector of this country is " + agr + " &#37; from its GDP"
	} else {
		return ""
	}
}

func getIndustrial(country *repo.Country) string {
	if country.IndustrialSector != 0 {
		agr := strconv.FormatFloat(country.IndustrialSector, 'f', -1, 64)
		return "Industrial sector of this country is " + agr + " &#37; from its GDP"
	} else {
		return ""
	}
}

func getService(country *repo.Country) string {
	if country.ServiceSector != 0 {
		agr := strconv.FormatFloat(country.ServiceSector, 'f', -1, 64)
		return "Service sector of this country is " + agr + " &#37; from its GDP"
	} else {
		return ""
	}
}
