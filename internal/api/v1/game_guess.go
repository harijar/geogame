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
		a.logger.Debug("Client side error: missing country input", err)
		return
	}
	a.logger.Debugf("User's guess: %s", request.Guess)

	countryID, err := c.Cookie("country")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game has not started"})
		a.logger.Debug("Client side error: game has not started", err)
		return
	}
	countryIDi, err := strconv.Atoi(countryID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid country id"})
		a.logger.Debug("Client side error: invalid country id", err)
		return
	}
	country := a.countries.Get(countryIDi)
	if country == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid country id"})
		a.logger.Debug("Client side error: invalid country id, could not get it from database")
		return
	}

	promptsStr, err := c.Cookie("prompts")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game has not started"})
		a.logger.Debug("Client side error: game has not started", err)
		return
	}
	prev := make([]*prompts.Prompt, 0)
	err = json.Unmarshal([]byte(promptsStr), &prev)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid prompts"})
		a.logger.Debug("Client side error: invalid prompts in cookie", err)
		return
	}

	response := GuessResponse{}
	for _, alias := range country.Aliases {
		if levenshtein.ComputeDistance(request.Guess, alias) <= 1 {
			response.Right = true
			response.Country = country.Name
			a.setCookie(c, "prompts", "", true)
			c.JSON(200, &response)
			a.logger.Debugf("User guessed the country in %v tries", len(prev)+1)
			return
		}
	}
	response.Right = false

	if a.triesLimit == len(prev) {
		response.Country = country.Name
		a.setCookie(c, "prompts", "", true)
		a.logger.Debugf("User didn't guess the country in %v tries", a.triesLimit)
	} else {
		prompt, err := a.prompts.GenRandom(country, prev)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error(err)
			return
		}
		prev = append(prev, prompt)
		prevOut, err := json.Marshal(&prev)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			a.logger.Error(err)
			return
		}
		a.setCookie(c, "prompts", string(prevOut), false)
		a.logger.Debugf("Next prompt: %s", prompt.Text)
		response.Prompt = prompt.Text
	}
	c.JSON(200, &response)
}
