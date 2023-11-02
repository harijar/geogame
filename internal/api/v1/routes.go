package v1

func (a *V1) registerRoutes() {
	a.router.POST("/game/start", a.gameStart)
	a.router.POST("/game/guess", a.gameGuess)
}
