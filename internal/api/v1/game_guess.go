package v1

import (
	"encoding/json"
	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/service/prompts"
	"net/http"
	"strconv"
)

type GuessRequest struct {
	Guess string `json:"guess"`
}

type GuessResponse struct {
	Right   bool   `json:"right"`
	Country string `json:"country"`
	Prompt  string `json:"prompt"`
}

func (a *V1) gameGuess(c *gin.Context) {
	request := GuessRequest{}
	err := c.BindJSON(&request)
	if err != nil || request.Guess == "" {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "missing country input"})
		return
	}

	countryID, err := c.Cookie("country")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game has not started"})
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

	promptsStr, err := c.Cookie("prompts")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game has not started"})
		return
	}
	prev := make([]*prompts.Prompt, 0)
	err = json.Unmarshal([]byte(promptsStr), &prev)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid prompts"})
		return
	}

	response := GuessResponse{}
	for _, alias := range country.Aliases {
		if levenshtein.ComputeDistance(request.Guess, alias) <= 1 {
			response.Right = true
			response.Country = country.Name
			a.setCookie(c, "prompts", "", true)
			c.JSON(200, &response)
			return
		}
	}
	response.Right = false

	if a.triesLimit == len(prev) {
		response.Country = country.Name
		a.setCookie(c, "prompts", "", true)
	} else {
		prompt, err := a.prompts.GenRandom(country, prev)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			return
		}
		prev = append(prev, prompt)
		prevOut, err := json.Marshal(&prev)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			return
		}
		a.setCookie(c, "prompts", string(prevOut), false)
		response.Prompt = prompt.Text
	}
	c.JSON(200, &response)
}
