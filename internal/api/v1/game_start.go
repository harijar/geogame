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
	a.logger.Debug("Game started")
	prevCountry, _ := c.Cookie("country")
	var country *countries.Country
	for country == nil {
		country = a.countries.GetRandom()
		if strconv.FormatInt(int64(country.ID), 10) == prevCountry {
			country = nil
		}
	}
	a.logger.Debugf("Current country is %s with ID %v", country.Name, country.ID)

	prompt, err := a.prompts.GenRandom(country, []*prompts.Prompt{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error(err)
		return
	}
	a.logger.Debugf("First prompt: %s", prompt.Text)

	promptsOut, err := json.Marshal([]*prompts.Prompt{prompt})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error(err)
		return
	}
	a.setCookie(c, "country", strconv.Itoa(country.ID), false)
	a.setCookie(c, "prompts", string(promptsOut), false)
	c.JSON(200, &StartResponse{prompt.Text})
}
