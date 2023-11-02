package ui

import "github.com/gin-gonic/gin"

func (ui *UI) GetGame(c *gin.Context) {
	// there will be html page with information about the game and button redirecting to "/game/play"
}

func (ui *UI) GetPlay(c *gin.Context) {
	// display game start html page

	ui.country = ui.countries.GetRandom()
	id, first, err := ui.prompt.GenRandom(ui.country, []int{}) // how to return error here? this function signature does not allow returning errors
	ui.prompts = make([]int, 0, ui.promptsLimit)
	ui.prompts = append(ui.prompts, id)

	// display first in the html page
}

func (ui *UI) PostPlay(c *gin.Context) {
	countryGot := c.PostForm("country")
	if countryGot == ui.country.Name {
		// display "you guessed the country" with "play again?" and "to the main page" buttons
	}
	if len(ui.prompts) == ui.promptsLimit {
		// display "the country was <ui.country.Name>" with "play again?" and "to the main page" buttons
	}
	id, next, err := ui.prompt.GenRandom(ui.country, ui.prompts)
	ui.prompts = append(ui.prompts, id)

	// display next in the html page
}
