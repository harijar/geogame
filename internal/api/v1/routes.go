package v1

func (a *V1) registerRoutes() {
	a.server.POST("/game/start", a.gameStart)
	a.server.POST("/game/guess", a.gameGuess)
}
