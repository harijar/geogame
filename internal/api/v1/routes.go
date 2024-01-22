package v1

func (a *V1) registerRoutes() {
	a.server.GET("v1/ws", a.serveWS)
	a.server.POST("v1/game/start", a.gameStart)
	a.server.POST("v1/game/guess", a.gameGuess)
	a.server.POST("v1/auth", a.auth)
	a.server.GET("v1/auth", a.authCheck)
	a.server.GET("v1/profile", a.profile)
	a.server.POST("/v1/profile/settings", a.updateProfileSettings)
	a.server.GET("/v1/profile/settings", a.getProfileSettings)
}
