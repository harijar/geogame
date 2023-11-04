package v1

import (
	"encoding/json"
	"github.com/agnivade/levenshtein"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/service/prompts"
	"net/http"
	"strconv"
	"strings"
)

type GuessResponse struct {
	Right   bool   `json:"right"`
	Country string `json:"country"`
	Prompt  string `json:"prompt"`
}

func (a *V1) gameGuess(c *gin.Context) {
	countryGot := strings.ToLower(c.PostForm("country"))
	if countryGot == "" {
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
		c.AbortWithStatusJSON(http.StatusNotFound, &gin.H{"error": "game has not started"})
		return
	}
	shown := make([]*prompts.Prompt, 0)
	err = json.Unmarshal([]byte(promptsStr), &shown)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid prompts"})
		return
	}

	if a.triesLimit == len(shown) {
		response.Country = country.Name
	} else {
		prompt, err := a.prompts.GenRandom(country, shown)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			return
		}
		shown = append(shown, prompt)
		promptsOut, err := json.Marshal(&shown)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
			return
		}
		c.SetCookie("prompts", string(promptsOut), 0, "/", c.Request.Host, false, true)
		response.Prompt = prompt.Text
	}
	c.JSON(200, &response)
}
