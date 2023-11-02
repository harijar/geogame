package api

import (
	"encoding/json"
	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/countries"
	"net/http"
	"strconv"
	"strings"
)

type StartResponse struct {
	Prompt string `json:"prompt"`
}

type GuessResponse struct {
	Right   bool   `json:"json:right,omitempty"`
	Country string `json:"country,omitempty"`
	Prompt  string `json:"prompt,omitempty"`
}

func (a *API) Start(c *gin.Context) {
	prevCountry, _ := c.Cookie("country")
	var country *countries.Country
	for {
		country = a.countries.GetRandom()
		prevCountryID, _ := strconv.Atoi(prevCountry)
		if country.ID != prevCountryID {
			break
		}
	}
	prompt, _, err := a.prompt.GenRandom(country, []int{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		return
	}
	c.JSON(200, &StartResponse{strconv.Itoa(prompt)})
}

func (a *API) Guess(c *gin.Context) {
	countryGot := strings.ToLower(c.PostForm("country"))
	if countryGot == "" {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "missing country input"})
		return
	}

	countryID, err := c.Cookie("country")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game not started"})
		return
	}
	countryIDi, err := strconv.Atoi(countryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid country id"})
		return
	}
	country := a.countries.Get(countryIDi)
	if country == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid country id"})
		return
	}

	response := GuessResponse{}
	for _, alias := range country.Aliases {
		if levenshtein.ComputeDistance(countryGot, alias) <= 1 {
			response.Right = true
			response.Country = country.Name
			c.JSON(200, &response)
			return
		}
	}
	response.Right = false

	promptsStr, err := c.Cookie("prompts")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game not started"})
		return
	}
	prompts := make([]int, 0)
	err = json.Unmarshal([]byte(promptsStr), &prompts)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid country id"})
		return
	}

	if a.triesLimit == len(prompts) {
		response.Country = country.Name
	} else {
		id, prompt, err := a.prompt.GenRandom(country, prompts)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			return
		}
		prompts = append(prompts, id)
		promptsOut, err := json.Marshal(&prompts)
		c.SetCookie("prompts", string(promptsOut), -1, "/", "localhost", false, true)
		response.Prompt = prompt
	}
	c.JSON(200, &response)
}
