package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (a *API) Start(c *gin.Context) {
	for _, cookie := range c.Request.Cookies() {
		cookie.Expires = time.Now()
		c.Request.AddCookie(cookie)
	}
	country := a.countries.GetRandom()
	c.Set("country", country.Name)
	// request to frontend
}

func (a *API) Play(c *gin.Context) {
	// display start page
	countryGot := c.PostForm("country")
	if countryGot == c.GetString("country") {
		// request to frontend
	}

	prompts := make([]int, 0, a.triesLimit)
	for _, cookie := range c.Request.Cookies() {
		id, err := strconv.Atoi(cookie.Value)
		if err != nil {
			// send error to frontend?
		}
		prompts = append(prompts, id)
	}

	if a.triesLimit == len(prompts) {
		// request to frontend
	}

	id, next, err := a.prompt.GenRandom(a.country, []int{})
	if err != nil {
		// send error to frontend?
	}
	c.SetCookie(strconv.Itoa(id), "", -1, "/", "localhost", false, true)
	// send next to frontend
}
