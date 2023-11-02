package ui

import "github.com/gin-gonic/gin"

func (ui *UI) registerRoutes(r *gin.Engine) {
	r.GET("/game", ui.GetGame)
	r.GET("/game/play", ui.GetPlay)
	r.POST("/game/play", ui.PostPlay)
}
