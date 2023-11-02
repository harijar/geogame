package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/countries"
	"net/http"
	"strconv"
)

type StartResponse struct {
	Prompt int `json:"prompt"`
}

type GuessResponse struct {
	Right   bool   `json:"json:right,omitempty"`
	Country string `json:"country,omitempty"`
	Prompt  string `json:"prompt,omitempty"`
}

func (v *V1) gameStart(c *gin.Context) {
	prevCountry, _ := c.Cookie("country")
	var country *countries.Country
	for country == nil {
		country = v.countries.GetRandom()
		if strconv.FormatInt(int64(country.ID), 10) == prevCountry {
			country = nil
		}
	}
	prompt, _, err := v.prompts.GenRandom(country, []int{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		return
	}

	promptsOut, err := json.Marshal([]int{prompt})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, &gin.H{"error": "invalid prompt id"})
		return
	}
	c.SetCookie("country", strconv.Itoa(country.ID), -1, "/", "localhost", false, true)
	c.SetCookie("prompts", string(promptsOut), -1, "/", "localhost", false, true)
	c.JSON(200, &StartResponse{prompt})
}
