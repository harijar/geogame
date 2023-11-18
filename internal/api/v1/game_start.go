package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/harijar/geogame/internal/repo/countries"
	"github.com/harijar/geogame/internal/service/prompts"
	"go.uber.org/zap"
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
	a.logger.Debug("current country info",
		zap.String("countryName", country.Name),
		zap.Int("countryID", country.ID))

	prompt, err := a.prompts.GenRandom(country, []*prompts.Prompt{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("prompt generation error", zap.Error(err))
		return
	}
	a.logger.Debug("first prompt",
		zap.String("promptText", prompt.Text),
		zap.Int("tryNumber", 1))

	promptsOut, err := json.Marshal([]*prompts.Prompt{prompt})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{"error": "internal server error"})
		a.logger.Error("prompt json encoding error", zap.Error(err))
		return
	}
	a.setCookie(c, "country", strconv.Itoa(country.ID), false)
	a.setCookie(c, "prompts", string(promptsOut), false)
	c.JSON(200, &StartResponse{prompt.Text})
}
