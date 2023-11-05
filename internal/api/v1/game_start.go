package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
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
	prompt, err := a.prompts.GenRandom(country, []*prompts.Prompt{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		return
	}

	promptsOut, err := json.Marshal([]*prompts.Prompt{prompt})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		return
	}
	c.SetCookie("country", strconv.Itoa(country.ID), 0, "/", c.Request.Host, false, true)
	c.SetCookie("prompts", string(promptsOut), 0, "/", c.Request.Host, false, true)
	c.JSON(200, &StartResponse{prompt.Text})
}
