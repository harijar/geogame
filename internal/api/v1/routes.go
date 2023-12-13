package v1

func (a *V1) registerRoutes() {
	a.server.POST("v1/game/start", a.gameStart)
	a.server.POST("v1/game/guess", a.gameGuess)
	a.server.POST("v1/auth", a.auth)
	a.server.GET("v1/auth", a.checkAuth)
	a.server.GET("v1/profile", a.profile)
}
