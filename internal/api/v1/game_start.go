package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/countries"
	"net/http"
	"strconv"
)

type StartResponse struct {
	Prompt string `json:"prompt"`
}

func (a *V1) gameStart(c *gin.Context) {
	prevCountry, _ := c.Cookie("country")
	var country *countries.Country
	for country == nil {
		country = a.countries.GetRandom()
		if strconv.FormatInt(int64(country.ID), 10) == prevCountry {
			country = nil
		}
	}
	id, prompt, err := a.prompts.GenRandom(country, []int{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		return
	}

	promptsOut, err := json.Marshal([]int{id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "invalid prompts"})
		return
	}
	c.SetCookie("country", strconv.Itoa(country.ID), -1, "/", c.Request.Host, false, true)
	c.SetCookie("prompts", string(promptsOut), -1, "/", c.Request.Host, false, true)
	c.JSON(200, &StartResponse{prompt})
}
