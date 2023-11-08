package v1

func (a *V1) registerRoutes() {
	a.server.POST("v1/game/start", a.gameStart)
	a.server.POST("v1/game/guess", a.gameGuess)
}
